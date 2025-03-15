package dto

type AuthRefreshTokenInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
