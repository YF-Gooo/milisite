package cache

import (
	"fmt"
	"strconv"
)

const (
	// DailyRankKey 每日排行
	DailyRankKey = "rank:daily"
)

// ImageViewKey 视频点击数的key
// view:image:1 -> 100
// view:image:2 -> 150
func ImageViewKey(id uint) string {
	return fmt.Sprintf("view:image:%s", strconv.Itoa(int(id)))
}
