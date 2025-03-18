package media

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"math/rand"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/creasty/defaults"
	"github.com/disintegration/imaging"
	"github.com/irvanherz/gourze/config"
	"github.com/irvanherz/gourze/modules/media/dto"
	"github.com/jinzhu/copier"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ImageSize struct {
	ID     string `json:"id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

var photoImageSizes = []ImageSize{
	{"sm", 150, 150},
	{"md", 300, 300},
	{"lg", 600, 600},
}

type MediaService interface {
	FindManyMedia(filter *dto.MediaFilterInput) ([]Media, int64, error)
	FindMediaByID(id uint) (*Media, error)
	UpdateMediaByID(id uint, input *dto.MediaUpdateInput) (*Media, error)
	DeleteMediaByID(id uint) (*Media, error)
	UploadPhoto(file multipart.File, originalName string) (*Media, error)
	UploadVideoViaTus(input *dto.MediaUploadVideoViaTusInput) (*dto.MediaUploadVideoViaTusResult, error)
}

type mediaService struct {
	Db           *gorm.DB
	Config       *config.Config
	BunnyService BunnyService
}

func NewMediaService(db *gorm.DB, conf *config.Config, bunnyService BunnyService) MediaService {
	return &mediaService{Db: db, Config: conf, BunnyService: bunnyService}
}

func (s *mediaService) FindManyMedia(filter *dto.MediaFilterInput) ([]Media, int64, error) {
	var medias []Media
	var count int64

	if err := defaults.Set(filter); err != nil {
		return nil, 0, err
	}
	query := s.Db
	query = filter.ApplyFilter(query)

	if err := query.Model(&Media{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	query = filter.ApplyPagination(query)

	if err := query.Find(&medias).Error; err != nil {
		return nil, 0, err
	}
	return medias, count, nil
}

func (s *mediaService) FindMediaByID(id uint) (*Media, error) {
	var media Media
	if err := s.Db.First(&media, id).Error; err != nil {
		return nil, err
	}
	return &media, nil
}

func (s *mediaService) UpdateMediaByID(id uint, input *dto.MediaUpdateInput) (*Media, error) {
	var media Media
	if err := s.Db.First(&media, id).Error; err != nil {
		return nil, err
	}
	copier.Copy(&media, &input)
	if err := s.Db.Save(&media).Error; err != nil {
		return nil, err
	}
	return &media, nil
}

func (s *mediaService) DeleteMediaByID(id uint) (*Media, error) {
	var media Media
	if err := s.Db.First(&media, id).Error; err != nil {
		return nil, err
	}
	if err := s.Db.Delete(&Media{}, id).Error; err != nil {
		return nil, err
	}
	return &media, nil
}

func (s *mediaService) UploadPhoto(file multipart.File, originalFileName string) (*Media, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	imageFiles := make([]ImageMediaFileData, len(photoImageSizes))
	originalFileNameNoExt := strings.TrimSuffix(originalFileName, filepath.Ext(originalFileName))
	randomizedFileName := generateRandomNumericFilename()

	for i, sz := range photoImageSizes {
		// Resize image
		resizedImage := imaging.Resize(img, sz.Width, sz.Height, imaging.Lanczos)
		resizedFileName := fmt.Sprintf("%s_%s.jpg", randomizedFileName, sz.ID)

		// Explicitly close before uploading
		resizedBuffer := new(bytes.Buffer)
		if err := jpeg.Encode(resizedBuffer, resizedImage, nil); err != nil {
			return nil, fmt.Errorf("failed to encode JPEG: %w", err)
		}

		// Upload to BunnyCDN
		downloadUrl, err := s.BunnyService.UploadFile(resizedBuffer.Bytes(), "/"+resizedFileName)
		if err != nil {
			return nil, fmt.Errorf("failed to upload to BunnyCDN: %w", err)
		}

		imageFiles[i] = ImageMediaFileData{
			ID:       sz.ID,
			Width:    sz.Width,
			Height:   sz.Height,
			FileName: resizedFileName,
			URL:      downloadUrl,
		}
	}

	var media Media
	mediaData := ImageMediaData{
		Version: "1.0",
		Files:   imageFiles,
	}
	mediaDataJson, err := json.Marshal(mediaData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}
	media.Title = originalFileNameNoExt
	media.Description = ""
	media.Type = Image
	media.Data = datatypes.JSON(mediaDataJson)

	// Save to DB
	if err := s.Db.Create(&media).Error; err != nil {
		return nil, err
	}

	return &media, nil
}

// UploadVideoViaTus implements MediaService.
func (s *mediaService) UploadVideoViaTus(input *dto.MediaUploadVideoViaTusInput) (*dto.MediaUploadVideoViaTusResult, error) {
	createdVideo, err := s.BunnyService.CreateVideo(&dto.BunnyCreateVideoInput{
		LibraryID: s.Config.Bunny.StreamLibraryID,
		Title:     input.Title,
	})
	if err != nil {
		return nil, err
	}
	mediaData := VideoMediaData{
		Version:  "1.0",
		Provider: "bunny",
		VideoID:  createdVideo.Guid,
	}
	mediaDataJson, err := json.Marshal(mediaData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}
	media := Media{
		Type:         Video,
		UploadStatus: Uploading,
		Title:        input.Title,
		Data:         mediaDataJson,
	}
	if err := s.Db.Save(&media).Error; err != nil {
		return nil, err
	}

	uploadSignature := s.BunnyService.ComputeTusUploadSignature(s.Config.Bunny.StreamLibraryID, createdVideo.Guid, s.Config.Bunny.StreamExpirationTime)
	return &dto.MediaUploadVideoViaTusResult{
		Headers: dto.MediaUploadVideoViaTusResultHeaders{
			AuthorizationSignature: uploadSignature,
			AuthorizationExpire:    s.Config.Bunny.StreamExpirationTime,
			VideoId:                createdVideo.Guid,
			LibraryId:              s.Config.Bunny.StreamLibraryID,
		},
		Metadata: dto.MediaUploadVideoViaTusResultMetadata{
			Filetype: input.Filetype,
			Title:    input.Title,
			UserID:   input.UserID,
		},
	}, nil
}

type VideoMediaData struct {
	Version  string `json:"version"`
	Provider string `json:"provider"`
	VideoID  string `json:"videoId"`
}

type ImageMediaData struct {
	Version string               `json:"version"`
	Files   []ImageMediaFileData `json:"files"`
}

type ImageMediaFileData struct {
	ID       string `json:"id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileName string `json:"fileName"`
	URL      string `json:"url"`
}

func generateRandomNumericFilename() string {
	timestamp := time.Now().UnixMicro()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := r.Intn(900000) + 100000
	return fmt.Sprintf("%d%d", timestamp, randomNumber)
}
