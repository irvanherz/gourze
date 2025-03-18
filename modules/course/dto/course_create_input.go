package dto

type CourseCreateInput struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  uint    `json:"categoryId"`
	UserID      uint    `json:"userId"`
}
