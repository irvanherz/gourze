package order

import (
	"time"

	"github.com/irvanherz/gourze/modules/user"
)

type OrderStatus string

const (
	Unpaid   OrderStatus = "unpaid"
	Paid     OrderStatus = "paid"
	Canceled OrderStatus = "canceled"
)

type Order struct {
	ID        uint        `gorm:"primarykey" json:"id"`
	UserID    uint        `gorm:"type:integer" json:"user_id"`
	Amount    float64     `gorm:"type:decimal(10,2)" json:"amount"`
	Status    OrderStatus `json:"status" gorm:"type:order_status;default:'unpaid'"`
	CreatedAt time.Time   `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt time.Time   `gorm:"type:timestamp" json:"updatedAt"`
	User      user.User   `json:"user" gorm:"foreignKey:UserID"`
	Items     []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	OrderID   uint      `gorm:"type:integer" json:"order_id"`
	CourseID  uint      `gorm:"type:integer" json:"course_id"`
	Quantity  int       `gorm:"type:integer" json:"quantity"`
	Price     float64   `gorm:"type:decimal(10,2)" json:"price"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updatedAt"`
}
