package model

import (
	"miliste/cache"
	"strconv"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
)

// Image 照片模型
type Image struct {
	gorm.Model
	Title  string
	Info   string
	Avatar string
}

// AvatarURL 封面地址
func (image *Image) AvatarURL() string {
	signedGetURL, _ := OssBucket.SignURL(image.Avatar, oss.HTTPGet, 600)
	return signedGetURL
}

// View 点击数
func (image *Image) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.ImageViewKey(image.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 照片游览
func (image *Image) AddView() {
	// 增加照片点击数
	cache.RedisClient.Incr(cache.ImageViewKey(image.ID))
	// 增加排行点击数
	cache.RedisClient.ZIncrBy(cache.DailyRankKey, 1, strconv.Itoa(int(image.ID)))
}
