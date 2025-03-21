package media

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/irvanherz/gourze/modules/media/dto"
	"github.com/irvanherz/gourze/utils"
)

type MediaController interface {
	FindManyMedia(*gin.Context)
	FindMediaByID(*gin.Context)
	UpdateMediaByID(*gin.Context)
	DeleteMediaByID(*gin.Context)
	UploadPhoto(*gin.Context)
	UploadVideoViaTus(*gin.Context)
}

type mediaController struct {
	Service MediaService
}

func NewMediaController(service MediaService) MediaController {
	return &mediaController{service}
}

func (mc *mediaController) UploadPhoto(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": err.Error()})
		return
	}
	defer file.Close()

	filename := header.Filename
	media, err := mc.Service.UploadPhoto(file, filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "ok", "message": "File uploaded successfully", "data": media})
}
func (mc *mediaController) FindManyMedia(c *gin.Context) {
	var filter dto.MediaFilterInput
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": err.Error()})
		return
	}
	medias, count, err := mc.Service.FindManyMedia(&filter)
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
		"data":    medias,
		"meta": gin.H{
			"numItems": count,
			"page":     page,
			"numPages": numPages,
			"take":     take,
		},
	})
}

func (mc *mediaController) FindMediaByID(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": "Invalid media ID"})
		return
	}
	media, err := mc.Service.FindMediaByID(uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "ok", "message": "Success", "data": media})
}

func (mc *mediaController) UpdateMediaByID(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": "Invalid media ID"})
		return
	}

	var media dto.MediaUpdateInput
	if err := c.ShouldBindJSON(&media); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": err.Error()})
		return
	}
	updatedMedia, err := mc.Service.UpdateMediaByID(uint(uid), &media)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "ok", "message": "Media updated successfully", "data": updatedMedia})
}

func (mc *mediaController) DeleteMediaByID(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": "Invalid media ID"})
		return
	}
	media, err := mc.Service.DeleteMediaByID(uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"code": "ok", "message": "Media deleted successfully", "data": media})
}

// UploadVideoViaTus implements MediaController.
func (mc *mediaController) UploadVideoViaTus(c *gin.Context) {
	var input dto.MediaUploadVideoViaTusInput
	currentUser, _ := utils.GetCurrentUser(c)
	if input.UserID == 0 {
		input.UserID = currentUser.ID
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid-params", "message": err.Error()})
		return
	}
	result, err := mc.Service.UploadVideoViaTus(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "internal-server-error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "ok", "message": "Success", "data": result})
}
