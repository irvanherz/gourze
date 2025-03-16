package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/irvanherz/gourze/config"
	"github.com/irvanherz/gourze/modules/user"
)

type AuthMiddleware interface {
	Authenticate() gin.HandlerFunc
	Authorize(mandatory bool, allowedRoles ...user.UserRole) gin.HandlerFunc
}

type authMiddleware struct {
	Config *config.Config
}

func (m *authMiddleware) Authorize(mandatory bool, allowedRoles ...user.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")

		if !exists {
			if mandatory {
				fmt.Println("Unauthorized access")
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			c.Next() // Guest access allowed
			return
		}

		// Retrieve user role from claims
		claims := user.(jwt.MapClaims)
		userRole, ok := claims["role"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Check if user role is allowed
		for _, role := range allowedRoles {
			if userRole == string(role) {
				c.Next()
				return
			}
		}

		// Role not allowed
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}

// Optional authentication - proceeds even if auth fails
func (m *authMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := m.parseAccessTokenFromCookie(c)
		if err == nil {
			c.Set("user", claims)
		}
		c.Next()
	}
}

func (m *authMiddleware) parseAccessTokenFromCookie(c *gin.Context) (jwt.MapClaims, error) {
	tokenString, err := c.Cookie("accessToken")
	if err != nil {
		return nil, fmt.Errorf("no token found")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return m.Config.Auth.JWTSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims format")
	}

	return claims, nil
}

func NewAuthMiddleware(config *config.Config) AuthMiddleware {
	return &authMiddleware{config}
}
