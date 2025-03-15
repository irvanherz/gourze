package user

import (
	"github.com/irvanherz/gourze/modules/user/dto"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserService interface {
	FindManyUsers() ([]User, error)
	CreateUser(user *dto.UserCreateInput) (*User, error)
	FindUserByID(id uint) (*User, error)
	UpdateUserByID(id uint, user *dto.UserUpdateInput) (*User, error)
	DeleteUserByID(id uint) (*User, error)
}

type userService struct {
	Db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{Db: db}
}

func (s *userService) FindManyUsers() ([]User, error) {
	var users []User
	if err := s.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) CreateUser(input *dto.UserCreateInput) (*User, error) {
	var user User
	copier.Copy(&user, &input)

	if err := s.Db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) FindUserByID(id uint) (*User, error) {
	var user User
	if err := s.Db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) UpdateUserByID(id uint, input *dto.UserUpdateInput) (*User, error) {
	var user User
	if err := s.Db.First(&user, id).Error; err != nil {
		return nil, err
	}
	copier.Copy(&user, &input)
	if err := s.Db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) DeleteUserByID(id uint) (*User, error) {
	var user User
	if err := s.Db.First(&user, id).Error; err != nil {
		return nil, err
	}
	if err := s.Db.Delete(&User{}, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
