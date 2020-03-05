package serializer

import "miliste/model"

// Image 视频序列化器
type Image struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	URL       string `json:"url"`
	Avatar    string `json:"avatar"`
	View      uint64 `json:"view"`
	CreatedAt int64  `json:"created_at"`
}

// BuildImage 序列化视频
func BuildImage(item model.Image) Image {
	return Image{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		URL:       item.ImageURL(),
		Avatar:    item.AvatarURL(),
		View:      item.View(),
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildImages 序列化视频列表
func BuildImages(items []model.Image) (images []Image) {
	for _, item := range items {
		image := BuildImage(item)
		images = append(images, image)
	}
	return images
}
