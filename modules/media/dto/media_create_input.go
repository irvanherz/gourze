package dto

type MediaCreateInput struct {
	Medianame string `json:"medianame"`
	Email     string `json:"email"`
	FullName  string `json:"fullName"`
}
