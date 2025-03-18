package auth

import (
	"testing"

	"github.com/irvanherz/gourze/config"
	"github.com/irvanherz/gourze/modules/auth/dto"
	"github.com/irvanherz/gourze/modules/user"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type AuthServiceTestSuite struct {
	suite.Suite
	db      *gorm.DB
	config  *config.Config
	service AuthService
}

func (suite *AuthServiceTestSuite) SetupTest() {
	suite.db = setupTestDB()
	suite.config = &config.Config{
		Auth: config.AuthConfig{
			JWTSecret: "testsecret",
		},
	}
	suite.service = NewAuthService(suite.db, suite.config)
}

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&user.User{})
	return db
}

func (suite *AuthServiceTestSuite) TestSignin_Success() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	suite.db.Create(&user.User{Username: "john_doe", Email: "john@doe.com", Password: string(hashedPassword)})

	input := dto.AuthSigninInput{
		UsernameOrEmail: "john_doe",
		Password:        "password123",
	}

	result, err := suite.service.Signin(input)
	suite.NoError(err)
	suite.NotNil(result)
	suite.NotEmpty(result.AccessToken)
}

func (suite *AuthServiceTestSuite) TestSignin_InvalidPassword() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	suite.db.Create(&user.User{Username: "john_doe", Email: "john@doe.com", Password: string(hashedPassword)})

	input := dto.AuthSigninInput{
		UsernameOrEmail: "john_doe",
		Password:        "wrongpassword",
	}

	result, err := suite.service.Signin(input)
	suite.Error(err)
	suite.Nil(result)
}

func (suite *AuthServiceTestSuite) TestSignin_UserNotFound() {
	input := dto.AuthSigninInput{
		UsernameOrEmail: "nonexistent_user",
		Password:        "password123",
	}

	result, err := suite.service.Signin(input)
	suite.Error(err)
	suite.Nil(result)
}

func TestAuthServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AuthServiceTestSuite))
}
