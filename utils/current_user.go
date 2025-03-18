package utils

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/irvanherz/gourze/modules/user"
)

// CurrentUser represents the authenticated user extracted from JWT
type CurrentUser struct {
	ID   uint
	Role user.UserRole
}

// GetCurrentUser extracts user claims from Gin context and converts them to CurrentUser
func GetCurrentUser(c *gin.Context) (*CurrentUser, error) {
	currentUser, exists := c.Get("user")
	if !exists {
		return nil, errors.New("unauthorized: user not found")
	}

	// Type assertion for jwt.MapClaims
	claims, ok := currentUser.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("unauthorized: invalid token data")
	}

	// Extract and convert "sub" (User ID)
	var userID uint
	if subStr, ok := claims["sub"].(string); ok {
		// Convert string to uint
		id, err := strconv.ParseUint(subStr, 10, 64)
		if err != nil {
			return nil, errors.New("unauthorized: invalid user ID format")
		}
		userID = uint(id)
	} else if subFloat, ok := claims["sub"].(float64); ok {
		// If stored as float64 (common with JSON numbers)
		userID = uint(subFloat)
	} else {
		return nil, errors.New("unauthorized: user ID missing or invalid")
	}

	// Extract "aud" (User Role) as string
	userRole, _ := claims["aud"].(string)
	parsedUserRole, _ := user.ParseUserRole(userRole)
	return &CurrentUser{
		ID:   userID,
		Role: parsedUserRole,
	}, nil
}
