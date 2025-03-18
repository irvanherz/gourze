package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/irvanherz/gourze/modules/auth/dto"
	"github.com/irvanherz/gourze/modules/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Signin(input dto.AuthSigninInput) (*dto.AuthResultDto, error) {
	args := m.Called(input)
	return args.Get(0).(*dto.AuthResultDto), args.Error(1)
}

func (m *MockAuthService) Signup(input dto.AuthSignupInput) (*dto.AuthResultDto, error) {
	args := m.Called(input)
	return args.Get(0).(*dto.AuthResultDto), args.Error(1)
}

func (m *MockAuthService) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockAuthService) CompareHashAndPassword(hashedPassword, password string) error {
	args := m.Called(hashedPassword, password)
	return args.Error(0)
}

func (m *MockAuthService) GenerateAccessToken(user user.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}
func (m *MockAuthService) GenerateRefreshToken(user user.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func TestSignup(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockAuthService)
	controller := NewAuthController(mockService)

	router := gin.Default()
	router.POST("/signup", controller.Signup)

	t.Run("successful signup", func(t *testing.T) {
		input := dto.AuthSignupInput{
			Username: "testuser",
			Password: "password",
		}
		output := dto.AuthResultDto{
			AccessToken:  "access-token",
			RefreshToken: "refresh-token",
		}

		mockService.On("Signup", input).Return(output, nil)

		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		mockService.On("Signup", input).Return(&dto.AuthResultDto{}, assert.AnError)
		assert.Equal(t, http.StatusOK, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("invalid params", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer([]byte(`{}`)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("internal server error", func(t *testing.T) {
		input := dto.AuthSignupInput{
			Username: "testuser",
			Password: "password",
		}

		mockService.On("Signup", input).Return(dto.AuthResultDto{}, assert.AnError)

		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		mockService.AssertExpectations(t)
	})
}
