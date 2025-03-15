package user

import (
	"github.com/irvanherz/gourze/modules/user/dto"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserService interface {
	FindMany() ([]User, error)
	Create(user *dto.UserCreateInput) error
	FindByID(id uint) (*User, error)
	UpdateByID(user *dto.UserUpdateInput) error
	DeleteByID(id uint) error
}

type userService struct {
	Db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{Db: db}
}

func (s *userService) FindMany() ([]User, error) {
	var users []User
	if err := s.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) Create(input *dto.UserCreateInput) error {
	var user User
	copier.Copy(&user, &input)

	if err := s.Db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (s *userService) FindByID(id uint) (*User, error) {
	var user User
	if err := s.Db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) UpdateByID(user *dto.UserUpdateInput) error {
	if err := s.Db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (s *userService) DeleteByID(id uint) error {
	if err := s.Db.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}
