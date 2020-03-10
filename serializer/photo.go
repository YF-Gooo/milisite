package serializer

import "miliste/model"

// Photo 照片序列化器
type Photo struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	Image     string `json:"image"`
	View      uint64 `json:"view"`
	CreatedAt int64  `json:"created_at"`
}

// BuildPhoto 序列化照片
func BuildPhoto(item model.Photo) Photo {
	return Photo{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		Image:     item.ImageURL(),
		View:      item.View(),
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildPhotos 序列化照片列表
func BuildPhotos(items []model.Photo) (photos []Photo) {
	for _, item := range items {
		photo := BuildPhoto(item)
		photos = append(photos, photo)
	}
	return photos
}
