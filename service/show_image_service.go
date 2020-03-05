package service

import (
	"miliste/model"
	"miliste/serializer"
)

// ShowImageService 投稿详情的服务
type ShowImageService struct {
}

// Show 视频
func (service *ShowImageService) Show(id string) serializer.Response {
	var image model.Image
	err := model.DB.First(&image, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}

	//处理视频被观看的一系问题
	image.AddView()

	return serializer.Response{
		Data: serializer.BuildImage(image),
	}
}
