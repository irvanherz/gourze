package order

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/irvanherz/gourze/modules/order/dto"
)

type OrderController interface {
	FindManyOrders(*gin.Context)
	CreateOrder(*gin.Context)
	FindOrderByID(*gin.Context)
	UpdateOrderByID(*gin.Context)
	DeleteOrderByID(*gin.Context)
}

type orderController struct {
	Service OrderService
}

func NewOrderController(service OrderService) OrderController {
	return &orderController{service}
}
func (oc *orderController) FindManyOrders(c *gin.Context) {
	var filter dto.OrderFilterInput
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": err.Error()})
		return
	}
	orders, count, err := oc.Service.FindManyOrders(&filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	page := filter.Page
	take := filter.Take
	numPages := (count + int64(take) - 1) / int64(take)

	c.JSON(http.StatusOK, gin.H{
		"code":    "ok",
		"message": "Success",
		"data":    orders,
		"meta": gin.H{
			"numItems": count,
			"page":     page,
			"numPages": numPages,
			"take":     take,
		},
	})
}

func (oc *orderController) CreateOrder(c *gin.Context) {
	var orderInput dto.OrderCreateInput
	if err := c.ShouldBindJSON(&orderInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": err.Error()})
		return
	}
	order, err := oc.Service.CreateOrder(&orderInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"code": "ok", "message": "Order created successfully", "data": order})
}

func (oc *orderController) FindOrderByID(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": "Invalid order ID"})
		return
	}
	order, err := oc.Service.FindOrderByID(uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "ok", "message": "Success", "data": order})
}

func (oc *orderController) UpdateOrderByID(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": "Invalid order ID"})
		return
	}

	var orderInput dto.OrderUpdateInput
	if err := c.ShouldBindJSON(&orderInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": err.Error()})
		return
	}
	order, err := oc.Service.UpdateOrderByID(uint(uid), &orderInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "ok", "message": "Order updated successfully", "data": order})
}

func (oc *orderController) DeleteOrderByID(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": "Invalid order ID"})
		return
	}
	order, err := oc.Service.DeleteOrderByID(uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"code": "ok", "message": "Order deleted successfully", "data": order})
}
