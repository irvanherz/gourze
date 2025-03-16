package media

import (
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/creasty/defaults"
	"github.com/disintegration/imaging"
	"github.com/go-resty/resty/v2"
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
	UpdateMediaByID(id uint, media *dto.MediaUpdateInput) (*Media, error)
	DeleteMediaByID(id uint) (*Media, error)
	UploadPhoto(file multipart.File, originalName string) (*Media, error)
}

type mediaService struct {
	Db     *gorm.DB
	Config *config.Config
}

func NewMediaService(db *gorm.DB, conf *config.Config) MediaService {
	return &mediaService{Db: db, Config: conf}
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
	resizedTempDir := os.TempDir()
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
		resizedFilePath := filepath.Join(resizedTempDir, resizedFileName)

		// Create temp file
		outFile, err := os.Create(resizedFilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to create file: %w", err)
		}

		// Explicitly close before uploading
		err = jpeg.Encode(outFile, resizedImage, &jpeg.Options{Quality: 90})
		outFile.Close()
		if err != nil {
			return nil, fmt.Errorf("failed to encode JPEG: %w", err)
		}

		// Upload to BunnyCDN
		downloadUrl, err := s.uploadToBunnyCDN(resizedFilePath, resizedFileName)
		if err != nil {
			return nil, fmt.Errorf("failed to upload to BunnyCDN: %w", err)
		}

		// Remove temp file **only if upload is successful**
		_ = os.Remove(resizedFilePath)

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

func (s *mediaService) uploadToBunnyCDN(filePath, fileName string) (string, error) {
	client := resty.New()
	url := fmt.Sprintf("https://%s.storage.bunnycdn.com/%s/%s", s.Config.Bunny.Region, s.Config.Bunny.StorageZone, fileName)

	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Read file contents
	fileData, err := io.ReadAll(file)

	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	resp, err := client.R().
		SetHeader("AccessKey", s.Config.Bunny.AccessKey).
		SetHeader("Content-Type", "image/jpeg").
		SetBody(fileData).
		Put(url)

	if err != nil || resp.StatusCode() != 201 {
		return "", fmt.Errorf("upload failed: %v, response: %s", err, resp.String())
	}

	return fmt.Sprintf("%s/%s", s.Config.Bunny.DownloadBaseURL, fileName), nil
}

type UploadResult struct {
	FileName string `json:"fileName"`
	URL      string `json:"url"`
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
	// Get current timestamp in milliseconds
	timestamp := time.Now().UnixMicro()

	// Generate a random 6-digit number
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random 6-digit number (100000 - 999999)
	randomNumber := r.Intn(900000) + 100000

	// Concatenate timestamp and random number
	return fmt.Sprintf("%d%d", timestamp, randomNumber)
}
