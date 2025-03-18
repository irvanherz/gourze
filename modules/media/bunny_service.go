package media

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/irvanherz/gourze/config"
	"github.com/irvanherz/gourze/modules/media/dto"
)

type BunnyService interface {
	UploadFile(data []byte, uploadPath string) (string, error)
	ComputeTusUploadSignature(libraryId uint64, videoId string, expirationTime uint64) string
	CreateVideo(input *dto.BunnyCreateVideoInput) (*dto.BunnyCreateVideoResponse, error)
}

type bunnyService struct {
	Config *config.Config
}

func NewBunnyService(conf *config.Config) BunnyService {
	return &bunnyService{Config: conf}
}

func (s *bunnyService) UploadFile(data []byte, uploadPath string) (string, error) {
	client := resty.New()
	url := fmt.Sprintf("https://%s.storage.bunnycdn.com/%s%s", s.Config.Bunny.Region, s.Config.Bunny.StorageZone, uploadPath)
	resp, err := client.R().
		SetHeader("AccessKey", s.Config.Bunny.AccessKey).
		SetHeader("Content-Type", "image/jpeg").
		SetBody(data).
		Put(url)

	if err != nil || resp.StatusCode() != 201 {
		return "", fmt.Errorf("upload failed: %v, response: %s", err, resp.String())
	}

	return fmt.Sprintf("%s/%s", s.Config.Bunny.DownloadBaseURL, uploadPath), nil
}

func (s *bunnyService) ComputeTusUploadSignature(libraryId uint64, videoId string, expirationTime uint64) string {
	hashData := fmt.Sprintf("%d%s%d%s", s.Config.Bunny.StreamLibraryID, s.Config.Bunny.StreamAccessKey, expirationTime, videoId)
	hash := sha256.Sum256([]byte(hashData))
	signatureString := hex.EncodeToString(hash[:])
	return signatureString
}

func (s *bunnyService) CreateVideo(input *dto.BunnyCreateVideoInput) (*dto.BunnyCreateVideoResponse, error) {
	client := resty.New()
	url := fmt.Sprintf("https://video.bunnycdn.com/library/%d/videos", input.LibraryID)
	resp, err := client.R().
		SetHeader("AccessKey", s.Config.Bunny.StreamAccessKey).
		SetHeader("Content-Type", "application/json").
		SetBody(input).
		Post(url)

	if err != nil {
		return nil, fmt.Errorf("failed to create video: %v", err)
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to create video, response: %s", resp.String())
	}

	var createVideoResponse dto.BunnyCreateVideoResponse
	if err := json.Unmarshal(resp.Body(), &createVideoResponse); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return &createVideoResponse, nil
}
