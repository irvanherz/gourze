package dto

type OrderStatus string

type OrderCreateInput struct {
	UserID uint                   `json:"user_id"`
	Items  []OrderItemCreateInput `json:"items"`
}

type OrderItemCreateInput struct {
	CourseID uint `json:"course_id"`
}
