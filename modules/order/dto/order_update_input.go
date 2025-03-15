package dto

type OrderUpdateInput struct {
	Ordername string `json:"ordername"`
	Email     string `json:"email"`
	FullName  string `json:"fullName"`
}
