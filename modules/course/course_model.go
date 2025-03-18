package course

import (
	"time"

	"github.com/irvanherz/gourze/modules/media"
	"github.com/irvanherz/gourze/modules/user"
	"gorm.io/datatypes"
)

type Course struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"type:varchar(100)" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	Price       float64        `gorm:"type:decimal(10,2)" json:"price"`
	CategoryID  uint           `gorm:"type:integer" json:"categoryId"`
	UserID      uint           `gorm:"type:integer" json:"userId"`
	Meta        datatypes.JSON `gorm:"type:jsonb;not null;default:'{}'" json:"meta"`
	CreatedAt   time.Time      `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"type:timestamp" json:"updatedAt"`
	User        user.User      `json:"user" gorm:"foreignKey:UserID"`
	Category    Category       `json:"category" gorm:"foreignKey:CategoryID"`
	Chapters    []Chapter      `json:"chapters" gorm:"foreignKey:CourseID"`
}

// Category model
type Category struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Name        string    `gorm:"type:varchar(100)" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"type:timestamp" json:"updatedAt"`
}

// Chapter model
type Chapter struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CourseID    uint           `gorm:"type:integer" json:"courseId"`
	Name        string         `gorm:"type:varchar(100)" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	Position    uint           `gorm:"type:integer" json:"position"`
	Duration    uint           `gorm:"type:integer" json:"durration"`
	MediaID     int            `gorm:"type:integer" json:"mediaId"`
	Meta        datatypes.JSON `gorm:"type:jsonb;not null;default:'{}'" json:"meta"`
	CreatedAt   time.Time      `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"type:timestamp" json:"updatedAt"`
	Media       media.Media    `json:"media" gorm:"foreignKey:MediaID"`
}

type CourseUser struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UserID    uint      `gorm:"type:integer" json:"userId"`
	CourseID  uint      `gorm:"type:integer" json:"courseId"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updatedAt"`
	User      user.User `json:"user" gorm:"foreignKey:UserID"`
	Course    Course    `json:"course" gorm:"foreignKey:CourseID"`
}

type CourseMeta struct {
	Provider     string `json:"provider" default:"bunny"`
	CollectionID string `json:"collectionId"`
}

type ChapterMeta struct {
	Provider string `json:"provider"`
	VideoID  string `json:"videoId"`
}
