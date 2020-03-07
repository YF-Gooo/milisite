package model

import (
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	OssClient *oss.Client
	OssBucket *oss.Bucket
)

func OssInit() {
	OssClient, _ = oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	OssBucket, _ = OssClient.Bucket(os.Getenv("OSS_BUCKET"))
}
