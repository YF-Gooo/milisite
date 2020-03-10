package api

import (
	service "miliste/service/photo"

	"github.com/gin-gonic/gin"
)

// CreatePhoto 照片投稿
func CreatePhoto(c *gin.Context) {
	service := service.CreatePhotoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowPhoto 照片详情接口
func ShowPhoto(c *gin.Context) {
	service := service.ShowPhotoService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// ListPhoto 照片列表接口
func ListPhoto(c *gin.Context) {
	service := service.ListPhotoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdatePhoto 更新照片的接口
func UpdatePhoto(c *gin.Context) {
	service := service.UpdatePhotoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeletePhoto 删除照片的接口
func DeletePhoto(c *gin.Context) {
	service := service.DeletePhotoService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}
