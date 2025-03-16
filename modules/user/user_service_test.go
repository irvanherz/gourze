package user

import (
	"testing"

	"github.com/irvanherz/gourze/modules/user/dto"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserServiceTestSuite struct {
	suite.Suite
	db      *gorm.DB
	service UserService
}

func (suite *UserServiceTestSuite) SetupTest() {
	suite.db = setupTestDB()
	suite.service = NewUserService(suite.db)
}

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&User{})
	return db
}

func (suite *UserServiceTestSuite) TestFindManyUsers() {
	// Seed data
	suite.db.Create(&User{Username: "john_doe", Email: "john@doe.com", FullName: "John Doe"})
	suite.db.Create(&User{Username: "jane_doe", Email: "jane@doe.com", FullName: "Jane Doe"})

	filter := &dto.UserFilterInput{
		SortBy:    "username",
		SortOrder: "asc",
		Page:      1,
		Take:      10,
	}

	users, count, err := suite.service.FindManyUsers(filter)
	suite.NoError(err)
	suite.Len(users, 2)
	suite.Equal(int64(2), count)
}

func (suite *UserServiceTestSuite) TestCreateUser() {
	input := &dto.UserCreateInput{FullName: "John Doe"}
	user, err := suite.service.CreateUser(input)
	suite.NoError(err)
	suite.Equal("John Doe", user.FullName)
}

func (suite *UserServiceTestSuite) TestFindUserByID() {
	// Seed data
	suite.db.Create(&User{FullName: "John Doe"})

	user, err := suite.service.FindUserByID(1)
	suite.NoError(err)
	suite.Equal("John Doe", user.FullName)
}

func (suite *UserServiceTestSuite) TestUpdateUserByID() {
	// Seed data
	suite.db.Create(&User{FullName: "John Doe"})

	input := &dto.UserUpdateInput{FullName: "John Smith"}
	user, err := suite.service.UpdateUserByID(1, input)
	suite.NoError(err)
	suite.Equal("John Smith", user.FullName)
}

func (suite *UserServiceTestSuite) TestDeleteUserByID() {
	// Seed data
	suite.db.Create(&User{FullName: "John Doe"})

	user, err := suite.service.DeleteUserByID(1)
	suite.NoError(err)
	suite.Equal("John Doe", user.FullName)

	// Verify user is deleted
	var count int64
	suite.db.Model(&User{}).Count(&count)
	suite.Equal(int64(0), count)
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
