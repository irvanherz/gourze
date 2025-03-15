package order

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/irvanherz/gourze/modules/order/dto"
)

type OrderController struct {
	Service OrderService
}

func NewOrderController(service OrderService) *OrderController {
	return &OrderController{service}
}

func (uc *OrderController) FindMany(c *gin.Context) {
	orders, err := uc.Service.FindMany()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (uc *OrderController) Create(c *gin.Context) {
	var order dto.OrderCreateInput
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uc.Service.Create(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, order)
}

func (uc *OrderController) FindOne(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}
	order, err := uc.Service.FindByID(uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (uc *OrderController) UpdateByID(c *gin.Context) {
	var order dto.OrderUpdateInput
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uc.Service.UpdateByID(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (uc *OrderController) DeleteByID(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}
	if err := uc.Service.DeleteByID(uint(uid)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
