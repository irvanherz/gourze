package course

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/irvanherz/gourze/modules/course/dto"
)

type CourseController interface {
	FindManyCourses(*gin.Context)
	FindCourseByID(*gin.Context)
	CreateCourse(*gin.Context)
	UpdateCourseByID(*gin.Context)
	DeleteCourseByID(*gin.Context)
}

type courseController struct {
	Service CourseService
}

func NewCourseController(service CourseService) CourseController {
	return &courseController{service}
}

func (cc *courseController) FindManyCourses(c *gin.Context) {
	var filter dto.CourseFilterInput
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": err.Error()})
		return
	}
	courses, count, err := cc.Service.FindManyCourses(&filter)
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
		"data":    courses,
		"meta": gin.H{
			"numItems": count,
			"page":     page,
			"numPages": numPages,
			"take":     take,
		},
	})
}

func (cc *courseController) CreateCourse(c *gin.Context) {
	var input dto.CourseCreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": err.Error()})
		return
	}
	course, err := cc.Service.CreateCourse(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"code": "ok", "message": "Course created successfully", "data": course})
}

func (cc *courseController) FindCourseByID(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": "Invalid course ID"})
		return
	}
	course, err := cc.Service.FindCourseByID(uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "ok", "message": "Success", "data": course})
}

func (cc *courseController) UpdateCourseByID(c *gin.Context) {
	var input dto.CourseUpdateInput
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": "Invalid course ID"})
		return
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": err.Error()})
		return
	}
	course, err := cc.Service.UpdateCourseByID(uint(uid), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "ok", "message": "Course updated successfully", "data": course})
}

func (cc *courseController) DeleteCourseByID(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": "Invalid course ID"})
		return
	}
	course, err := cc.Service.DeleteCourseByID(uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"code": "ok", "message": "Course deleted successfully", "data": course})
}
