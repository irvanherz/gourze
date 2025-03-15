package dto

type UserCreateInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
}
