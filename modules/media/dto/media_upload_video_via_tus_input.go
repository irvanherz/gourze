package dto

type MediaUploadVideoViaTusInput struct {
	UserID   uint   `json:"userID"`
	Title    string `json:"title"`
	Filetype string `json:"filetype"`
}
