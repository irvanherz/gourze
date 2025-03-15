package user

import (
	"time"
)

type UserRole string

const (
	Super   UserRole = "super"
	Admin   UserRole = "admin"
	Generic UserRole = "generic"
)

type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Username  string    `gorm:"unique;type:varchar(255)" json:"username"`
	Email     string    `gorm:"unique;type:varchar(255)" json:"email"`
	FullName  string    `gorm:"type:varchar(255)" json:"fullName"`
	Password  string    `gorm:"type:varchar(255)" json:"-"`
	Role      UserRole  `json:"role" gorm:"type:user_role;default:'generic'"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updatedAt"`
}
