package service

import (
	"miliste/model"
	"miliste/serializer"
)

// UpdatePhotoService 更新照片的服务
type UpdatePhotoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"max=300"`
}

// Update 更新照片
func (service *UpdatePhotoService) Update(id string) serializer.Response {
	var photo model.Photo
	err := model.DB.First(&photo, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "照片不存在",
			Error:  err.Error(),
		}
	}

	photo.Title = service.Title
	photo.Info = service.Info
	err = model.DB.Save(&photo).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "照片保存失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildPhoto(photo),
	}
}
