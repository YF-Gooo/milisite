package service

import (
	"miliste/model"
	"miliste/serializer"
)

// DeleteImageService 删除投稿的服务
type DeleteImageService struct {
}

// Delete 删除照片
func (service *DeleteImageService) Delete(id string) serializer.Response {
	var image model.Image
	err := model.DB.First(&image, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "照片不存在",
			Error:  err.Error(),
		}
	}

	err = model.DB.Delete(&image).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "照片删除失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{}
}
