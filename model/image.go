package model

import (
	"miliste/cache"
	"os"
	"strconv"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
)

// Image 视频模型
type Image struct {
	gorm.Model
	Title  string
	Info   string
	URL    string
	Avatar string
}

// AvatarURL 封面地址
func (image *Image) AvatarURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(image.Avatar, oss.HTTPGet, 600)
	return signedGetURL
}

// ImageURL 视频地址
func (image *Image) ImageURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(image.URL, oss.HTTPGet, 600)
	return signedGetURL
}

// View 点击数
func (image *Image) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.ImageViewKey(image.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 视频游览
func (image *Image) AddView() {
	// 增加视频点击数
	cache.RedisClient.Incr(cache.ImageViewKey(image.ID))
	// 增加排行点击数
	cache.RedisClient.ZIncrBy(cache.DailyRankKey, 1, strconv.Itoa(int(image.ID)))
}
