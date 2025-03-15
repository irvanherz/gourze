package dto

type AuthSignupInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	FullName string `json:"fullName" binding:"required"`
	Password string `json:"password" binding:"required"`
}
