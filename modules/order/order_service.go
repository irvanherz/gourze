package order

import (
	"github.com/creasty/defaults"
	"github.com/irvanherz/gourze/modules/order/dto"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type OrderService interface {
	FindManyOrders(filter *dto.OrderFilterInput) ([]Order, int64, error)
	CreateOrder(order *dto.OrderCreateInput) (*Order, error)
	FindOrderByID(id uint) (*Order, error)
	UpdateOrderByID(id uint, order *dto.OrderUpdateInput) (*Order, error)
	DeleteOrderByID(id uint) (*Order, error)
}

type orderService struct {
	Db *gorm.DB
}

func NewOrderService(db *gorm.DB) OrderService {
	return &orderService{Db: db}
}
func (s *orderService) FindManyOrders(filter *dto.OrderFilterInput) ([]Order, int64, error) {
	var orders []Order
	var count int64

	if err := defaults.Set(filter); err != nil {
		return nil, 0, err
	}
	query := s.Db
	query = filter.ApplyFilter(query)

	if err := query.Model(&Order{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	query = filter.ApplyPagination(query)

	if err := query.Find(&orders).Error; err != nil {
		return nil, 0, err
	}
	return orders, count, nil
}

func (s *orderService) CreateOrder(input *dto.OrderCreateInput) (*Order, error) {
	var order Order
	copier.Copy(&order, &input)

	if err := s.Db.Create(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *orderService) FindOrderByID(id uint) (*Order, error) {
	var order Order
	if err := s.Db.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *orderService) UpdateOrderByID(id uint, input *dto.OrderUpdateInput) (*Order, error) {
	var order Order
	if err := s.Db.First(&order, id).Error; err != nil {
		return nil, err
	}
	copier.Copy(&order, &input)
	if err := s.Db.Save(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *orderService) DeleteOrderByID(id uint) (*Order, error) {
	var order Order
	if err := s.Db.First(&order, id).Error; err != nil {
		return nil, err
	}
	if err := s.Db.Delete(&Order{}, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
