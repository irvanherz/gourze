package dto

type CourseUpdateInput struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	CategoryID  *uint    `json:"categoryId,omitempty"`
}
