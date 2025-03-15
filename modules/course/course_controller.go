package course

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/irvanherz/gourze/modules/course/dto"
)

type CourseController struct {
	Service CourseService
}

func NewCourseController(service CourseService) *CourseController {
	return &CourseController{service}
}

func (cc *CourseController) FindManyCourse(c *gin.Context) {
	courses, err := cc.Service.FindManyCourse()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func (cc *CourseController) CreateCourse(c *gin.Context) {
	var input dto.CourseCreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	course, err := cc.Service.CreateCourse(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, course)
}

func (cc *CourseController) FindOneCourse(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}
	course, err := cc.Service.FindCourseByID(uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, course)
}

func (cc *CourseController) UpdateCourse(c *gin.Context) {
	var input dto.CourseUpdateInput
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	course, err := cc.Service.UpdateCourseByID(uint(uid), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, course)
}

func (cc *CourseController) DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}
	course, err := cc.Service.DeleteCourseByID(uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, course)
}
