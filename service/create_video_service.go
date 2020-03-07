package service

import (
	"miliste/model"
	"miliste/serializer"
)

// CreateImageService 照片投稿的服务
type CreateImageService struct {
	Title  string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Info   string `form:"info" json:"info" binding:"max=3000"`
	Avatar string `form:"avatar" json:"avatar"`
}

// Create 创建照片
func (service *CreateImageService) Create() serializer.Response {
	image := model.Image{
		Title:  service.Title,
		Info:   service.Info,
		Avatar: service.Avatar,
	}

	err := model.DB.Create(&image).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "照片保存失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildImage(image),
	}
}
