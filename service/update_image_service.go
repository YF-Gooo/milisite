package service

import (
	"miliste/model"
	"miliste/serializer"
)

// UpdateImageService 更新视频的服务
type UpdateImageService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"max=300"`
}

// Update 更新视频
func (service *UpdateImageService) Update(id string) serializer.Response {
	var image model.Image
	err := model.DB.First(&image, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}

	image.Title = service.Title
	image.Info = service.Info
	err = model.DB.Save(&image).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "视频保存失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildImage(image),
	}
}
