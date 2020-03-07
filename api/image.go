package api

import (
	"miliste/service"

	"github.com/gin-gonic/gin"
)

// CreateImage 照片投稿
func CreateImage(c *gin.Context) {
	service := service.CreateImageService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowImage 照片详情接口
func ShowImage(c *gin.Context) {
	service := service.ShowImageService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// ListImage 照片列表接口
func ListImage(c *gin.Context) {
	service := service.ListImageService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdateImage 更新照片的接口
func UpdateImage(c *gin.Context) {
	service := service.UpdateImageService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteImage 删除照片的接口
func DeleteImage(c *gin.Context) {
	service := service.DeleteImageService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}
