package user

import (
	"testing"

	"github.com/irvanherz/gourze/modules/user/dto"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&User{})
	return db
}

func TestFindManyUsers(t *testing.T) {
	db := setupTestDB()
	service := NewUserService(db)

	// Seed data
	db.Create(&User{FullName: "John Doe"})
	db.Create(&User{FullName: "Jane Doe"})

	users, err := service.FindManyUsers()
	assert.NoError(t, err)
	assert.Len(t, users, 2)
}

func TestCreateUser(t *testing.T) {
	db := setupTestDB()
	service := NewUserService(db)

	input := &dto.UserCreateInput{FullName: "John Doe"}
	user, err := service.CreateUser(input)
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", user.FullName)
}

func TestFindUserByID(t *testing.T) {
	db := setupTestDB()
	service := NewUserService(db)

	// Seed data
	db.Create(&User{FullName: "John Doe"})

	user, err := service.FindUserByID(1)
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", user.FullName)
}

func TestUpdateUserByID(t *testing.T) {
	db := setupTestDB()
	service := NewUserService(db)

	// Seed data
	db.Create(&User{FullName: "John Doe"})

	input := &dto.UserUpdateInput{FullName: "John Smith"}
	user, err := service.UpdateUserByID(1, input)
	assert.NoError(t, err)
	assert.Equal(t, "John Smith", user.FullName)
}

func TestDeleteUserByID(t *testing.T) {
	db := setupTestDB()
	service := NewUserService(db)

	// Seed data
	db.Create(&User{FullName: "John Doe"})

	user, err := service.DeleteUserByID(1)
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", user.FullName)

	// Verify user is deleted
	var count int64
	db.Model(&User{}).Count(&count)
	assert.Equal(t, int64(0), count)
}
