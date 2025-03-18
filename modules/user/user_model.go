package user

import (
	"errors"
	"time"

	"gorm.io/datatypes"
)

type UserRole string

const (
	Super   UserRole = "super"
	Admin   UserRole = "admin"
	Generic UserRole = "generic"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Username  string         `gorm:"unique;type:varchar(255)" json:"username"`
	Email     string         `gorm:"unique;type:varchar(255)" json:"email"`
	FullName  string         `gorm:"type:varchar(255)" json:"fullName"`
	Password  string         `gorm:"type:varchar(255)" json:"-"`
	Role      UserRole       `json:"role" gorm:"type:user_role;default:'generic'"`
	Meta      datatypes.JSON `gorm:"type:jsonb;not null;default:'{}'" json:"meta"`
	CreatedAt time.Time      `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"type:timestamp" json:"updatedAt"`
}

func ParseUserRole(roleStr string) (UserRole, error) {
	switch roleStr {
	case string(Super):
		return Super, nil
	case string(Admin):
		return Admin, nil
	case string(Generic):
		return Generic, nil
	default:
		return "", errors.New("invalid user role")
	}
}
