package order

import (
	"github.com/irvanherz/gourze/modules/order/dto"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type OrderService interface {
	FindMany() ([]Order, error)
	Create(order *dto.OrderCreateInput) error
	FindByID(id uint) (*Order, error)
	UpdateByID(order *dto.OrderUpdateInput) error
	DeleteByID(id uint) error
}

type orderService struct {
	Db *gorm.DB
}

func NewOrderService(db *gorm.DB) OrderService {
	return &orderService{Db: db}
}

func (s *orderService) FindMany() ([]Order, error) {
	var orders []Order
	if err := s.Db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *orderService) Create(input *dto.OrderCreateInput) error {
	var order Order
	copier.Copy(&order, &input)

	if err := s.Db.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func (s *orderService) FindByID(id uint) (*Order, error) {
	var order Order
	if err := s.Db.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *orderService) UpdateByID(order *dto.OrderUpdateInput) error {
	if err := s.Db.Save(order).Error; err != nil {
		return err
	}
	return nil
}

func (s *orderService) DeleteByID(id uint) error {
	if err := s.Db.Delete(&Order{}, id).Error; err != nil {
		return err
	}
	return nil
}
