package dto

type AuthRefreshTokenInput struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}
