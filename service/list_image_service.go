package service

import (
	"miliste/model"
	"miliste/serializer"
)

// ListImageService 照片列表服务
type ListImageService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// List 照片列表
func (service *ListImageService) List() serializer.Response {
	images := []model.Image{}
	total := 0

	if service.Limit == 0 {
		service.Limit = 6
	}

	if err := model.DB.Model(model.Image{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&images).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildImages(images), uint(total))
}
