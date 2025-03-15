package media

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/irvanherz/gourze/modules/media/dto"
)

type MediaController struct {
	Service MediaService
}

func NewMediaController(service MediaService) *MediaController {
	return &MediaController{service}
}

func (uc *MediaController) UploadPhoto(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	defer file.Close()

	filename := header.Filename
	media, err := uc.Service.UploadPhoto(file, filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "data": media})

}

func (uc *MediaController) FindMany(c *gin.Context) {
	medias, err := uc.Service.FindMany()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, medias)
}

func (uc *MediaController) Create(c *gin.Context) {
	var media dto.MediaCreateInput
	if err := c.ShouldBindJSON(&media); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := uc.Service.Create(&media); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, media)
}

func (uc *MediaController) FindOne(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid media ID"})
		return
	}
	media, err := uc.Service.FindByID(uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, media)
}

func (uc *MediaController) UpdateByID(c *gin.Context) {
	var media dto.MediaUpdateInput
	if err := c.ShouldBindJSON(&media); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := uc.Service.UpdateByID(&media); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, media)
}

func (uc *MediaController) DeleteByID(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid media ID"})
		return
	}
	if err := uc.Service.DeleteByID(uint(uid)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
