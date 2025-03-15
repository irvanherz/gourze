package course

import (
	"time"

	"github.com/irvanherz/gourze/modules/media"
	"github.com/irvanherz/gourze/modules/user"
)

// Course model
type Course struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Name        string    `gorm:"type:varchar(100)" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Price       float64   `gorm:"type:decimal(10,2)" json:"price"`
	CategoryID  uint      `gorm:"type:integer" json:"category_id"`
	UserID      uint      `gorm:"type:integer" json:"user_id"`
	User        user.User `json:"user" gorm:"foreignKey:UserID"`
	Category    Category  `json:"category" gorm:"foreignKey:CategoryID"`
	Chapters    []Chapter `json:"chapters" gorm:"foreignKey:CourseID"`
	CreatedAt   time.Time `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"type:timestamp" json:"updatedAt"`
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
	ID          uint        `gorm:"primarykey" json:"id"`
	CourseID    uint        `gorm:"type:integer" json:"course_id"`
	Name        string      `gorm:"type:varchar(100)" json:"name"`
	Description string      `gorm:"type:text" json:"description"`
	Position    int         `gorm:"type:integer" json:"position"`
	MediaID     int         `gorm:"type:integer" json:"media_id"`
	Video       media.Media `json:"video" gorm:"foreignKey:MediaID"`
	CreatedAt   time.Time   `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt   time.Time   `gorm:"type:timestamp" json:"updatedAt"`
}

type CourseUser struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UserID    uint      `gorm:"type:integer" json:"user_id"`
	CourseID  uint      `gorm:"type:integer" json:"course_id"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updatedAt"`
	User      user.User `json:"user" gorm:"foreignKey:UserID"`
	Course    Course    `json:"course" gorm:"foreignKey:CourseID"`
}
