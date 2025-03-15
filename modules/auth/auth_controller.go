package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irvanherz/gourze/modules/auth/dto"
)

type AuthController struct {
	Service AuthService
}

func NewAuthController(service AuthService) *AuthController {
	return &AuthController{service}
}

func (uc *AuthController) Signin(c *gin.Context) {
	var input dto.AuthSigninInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": err.Error()})
		return
	}
	result, err := uc.Service.Signin(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "ok", "message": "Signin successful", "data": result})
}

func (uc *AuthController) Signup(c *gin.Context) {
	var input dto.AuthSignupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": err.Error()})
		return
	}
	result, err := uc.Service.Signup(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.SetCookie("accessToken", result.AccessToken, 3600, "/", "", false, true)
	c.SetCookie("refreshToken", result.RefreshToken, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"code": "ok", "message": "Signup successful", "data": result})
}
