package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/irvanherz/gourze/modules/user/dto"
)

type UserController struct {
	Service UserService
}

func NewUserController(service UserService) *UserController {
	return &UserController{service}
}

func (uc *UserController) FindManyUsers(c *gin.Context) {
	var filter dto.UserFilterInput
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": err.Error()})
		return
	}
	users, err := uc.Service.FindManyUsers(&filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "ok", "message": "Success", "data": users})
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var userInput dto.UserCreateInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": err.Error()})
		return
	}
	user, err := uc.Service.CreateUser(&userInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"code": "ok", "message": "User created successfully", "data": user})
}

func (uc *UserController) FindUserByID(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": "Invalid user ID"})
		return
	}
	user, err := uc.Service.FindUserByID(uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "ok", "message": "Success", "data": user})
}

func (uc *UserController) UpdateUserByID(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": "Invalid user ID"})
		return
	}

	var userInput dto.UserUpdateInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": err.Error()})
		return
	}
	user, err := uc.Service.UpdateUserByID(uint(uid), &userInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "ok", "message": "User updated successfully", "data": user})
}

func (uc *UserController) DeleteUserByID(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "ok", "message": "Invalid user ID"})
		return
	}
	user, err := uc.Service.DeleteUserByID(uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"code": "ok", "message": "User deleted successfully", "data": user})
}
