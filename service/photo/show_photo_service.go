package service

import (
	"miliste/model"
	"miliste/serializer"
)

// ShowPhotoService 投稿详情的服务
type ShowPhotoService struct {
}

// Show 照片
func (service *ShowPhotoService) Show(id string) serializer.Response {
	var photo model.Photo
	err := model.DB.First(&photo, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "照片不存在",
			Error:  err.Error(),
		}
	}

	//处理照片被观看的一系问题
	photo.AddView()

	return serializer.Response{
		Data: serializer.BuildPhoto(photo),
	}
}
