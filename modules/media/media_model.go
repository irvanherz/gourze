package media

import (
	"time"

	"gorm.io/datatypes"
)

type MediaType string

const (
	Image    MediaType = "image"
	Document MediaType = "document"
	Video    MediaType = "video"
)

type Media struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Type        MediaType      `gorm:"type:media_type;not null" json:"type"`
	Data        datatypes.JSON `gorm:"type:jsonb;not null" json:"data"`
	Title       string         `gorm:"type:varchar(255);not null" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	CreatedAt   time.Time      `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"type:timestamp" json:"updatedAt"`
}
