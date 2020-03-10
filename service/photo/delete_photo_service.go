package service

import (
	"miliste/model"
	"miliste/serializer"
)

// DeletePhotoService 删除投稿的服务
type DeletePhotoService struct {
}

// Delete 删除照片
func (service *DeletePhotoService) Delete(id string) serializer.Response {
	var photo model.Photo
	err := model.DB.First(&photo, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "照片不存在",
			Error:  err.Error(),
		}
	}

	err = model.DB.Delete(&photo).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "照片删除失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{}
}
