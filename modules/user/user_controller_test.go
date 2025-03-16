package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/irvanherz/gourze/modules/user/dto"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type UserControllerTestSuite struct {
	suite.Suite
	DB         *gorm.DB
	Router     *gin.Engine
	Service    UserService
	Controller UserController
}

func (suite *UserControllerTestSuite) SetupTest() {
	suite.DB = setupTestDB()
	suite.Service = NewUserService(suite.DB)
	suite.Controller = NewUserController(suite.Service)
	suite.Router = gin.Default()
	suite.Router.PUT("/users/:id", suite.Controller.UpdateUserByID)

	// Seed data
	suite.DB.Create(&User{FullName: "John Doe"})
}

func (suite *UserControllerTestSuite) TestUpdateUserByID_Success() {
	userInput := dto.UserUpdateInput{FullName: "John Smith"}
	jsonValue, _ := json.Marshal(userInput)
	req, _ := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	suite.Router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code)
	var response map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &response)
	suite.Equal("ok", response["code"])
	suite.Equal("User updated successfully", response["message"])
	suite.Equal("John Smith", response["data"].(map[string]interface{})["FullName"])
}

func (suite *UserControllerTestSuite) TestUpdateUserByID_InvalidUserID() {
	userInput := dto.UserUpdateInput{FullName: "John Smith"}
	jsonValue, _ := json.Marshal(userInput)
	req, _ := http.NewRequest(http.MethodPut, "/users/abc", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	suite.Router.ServeHTTP(resp, req)

	suite.Equal(http.StatusBadRequest, resp.Code)
	var response map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &response)
	suite.Equal("invalid-params", response["code"])
	suite.Equal("Invalid user ID", response["message"])
}

func (suite *UserControllerTestSuite) TestUpdateUserByID_InvalidInputData() {
	invalidInput := map[string]interface{}{"FullName": 123}
	jsonValue, _ := json.Marshal(invalidInput)
	req, _ := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	suite.Router.ServeHTTP(resp, req)

	suite.Equal(http.StatusBadRequest, resp.Code)
	var response map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &response)
	suite.Equal("invalid-params", response["code"])
}

func (suite *UserControllerTestSuite) TestUpdateUserByID_UserNotFound() {
	userInput := dto.UserUpdateInput{FullName: "John Smith"}
	jsonValue, _ := json.Marshal(userInput)
	req, _ := http.NewRequest(http.MethodPut, "/users/999", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	suite.Router.ServeHTTP(resp, req)

	suite.Equal(http.StatusInternalServerError, resp.Code)
	var response map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &response)
	suite.Equal("internal-server-error", response["code"])
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
