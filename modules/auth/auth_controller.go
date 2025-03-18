package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irvanherz/gourze/modules/auth/dto"
)

type AuthController interface {
	Signin(*gin.Context)
	Signup(*gin.Context)
}

type authController struct {
	Service AuthService
}

func NewAuthController(service AuthService) AuthController {
	return &authController{service}
}

func (ac *authController) Signin(c *gin.Context) {
	var input dto.AuthSigninInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": err.Error()})
		return
	}
	result, err := ac.Service.Signin(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.SetCookie("accessToken", result.AccessToken, 3600, "/", "", false, true)
	c.SetCookie("refreshToken", result.RefreshToken, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"code": "ok", "message": "Signin successful", "data": result})
}

func (ac *authController) Signup(c *gin.Context) {
	var input dto.AuthSignupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": err.Error()})
		return
	}
	result, err := ac.Service.Signup(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.SetCookie("accessToken", result.AccessToken, 3600, "/", "", false, true)
	c.SetCookie("refreshToken", result.RefreshToken, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"code": "ok", "message": "Signup successful", "data": result})
}
