package model

import (
	"miliste/cache"
	"strconv"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
)

// Photo 照片模型
type Photo struct {
	gorm.Model
	Title string
	Info  string `gorm:"size:100000"`
	Image string
}

// ImageURL 图片地址
func (photo *Photo) ImageURL() string {
	signedGetURL, _ := OssBucket.SignURL(photo.Image, oss.HTTPGet, 600)
	return signedGetURL
}

// View 点击数
func (photo *Photo) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.PhotoViewKey(photo.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 照片游览
func (photo *Photo) AddView() {
	// 增加照片点击数
	cache.RedisClient.Incr(cache.PhotoViewKey(photo.ID))
	// 增加排行点击数
	cache.RedisClient.ZIncrBy(cache.DailyRankKey, 1, strconv.Itoa(int(photo.ID)))
}
