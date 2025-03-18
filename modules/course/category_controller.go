package course

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/irvanherz/gourze/modules/course/dto"
	"github.com/irvanherz/gourze/modules/user"
	"github.com/irvanherz/gourze/utils"
)

type CategoryController interface {
	FindManyCategories(*gin.Context)
	FindCategoryByID(*gin.Context)
	CreateCategory(*gin.Context)
	UpdateCategoryByID(*gin.Context)
	DeleteCategoryByID(*gin.Context)
}

type categoryController struct {
	Service CategoryService
}

func NewCategoryController(service CategoryService) CategoryController {
	return &categoryController{service}
}

func (cc *categoryController) FindManyCategories(c *gin.Context) {
	var filter dto.CategoryFilterInput
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": err.Error()})
		return
	}
	categories, count, err := cc.Service.FindManyCategories(&filter)
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
		"data":    categories,
		"meta": gin.H{
			"numItems": count,
			"page":     page,
			"numPages": numPages,
			"take":     take,
		},
	})
}

func (cc *categoryController) CreateCategory(c *gin.Context) {
	var input dto.CategoryCreateInput
	currentUser, _ := utils.GetCurrentUser(c)

	if currentUser.Role != user.Super && currentUser.Role != user.Admin {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": "unauthorized", "message": "Unauthorized"})
		return
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": err.Error()})
		return
	}
	category, err := cc.Service.CreateCategory(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"code": "ok", "message": "Category created successfully", "data": category})
}

func (cc *categoryController) FindCategoryByID(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": "Invalid category ID"})
		return
	}
	category, err := cc.Service.FindCategoryByID(uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "ok", "message": "Success", "data": category})
}

func (cc *categoryController) UpdateCategoryByID(c *gin.Context) {
	var input dto.CategoryUpdateInput
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": "Invalid category ID"})
		return
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": err.Error()})
		return
	}
	category, err := cc.Service.UpdateCategoryByID(uint(uid), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "ok", "message": "Category updated successfully", "data": category})
}

func (cc *categoryController) DeleteCategoryByID(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": "Invalid category ID"})
		return
	}
	category, err := cc.Service.DeleteCategoryByID(uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"code": "ok", "message": "Category deleted successfully", "data": category})
}
