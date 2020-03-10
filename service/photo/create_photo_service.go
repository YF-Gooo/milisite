package service

import (
	"miliste/model"
	"miliste/serializer"
)

// CreatePhotoService 照片投稿的服务
type CreatePhotoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Info  string `form:"info" json:"info" binding:"max=100000"`
	Image string `form:"image" json:"image"`
}

// Create 创建照片
func (service *CreatePhotoService) Create() serializer.Response {
	photo := model.Photo{
		Title: service.Title,
		Info:  service.Info,
		Image: service.Image,
	}

	err := model.DB.Create(&photo).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "照片保存失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildPhoto(photo),
	}
}
