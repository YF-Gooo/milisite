package cache

import (
	"fmt"
	"strconv"
)

const (
	// DailyRankKey 每日排行
	DailyRankKey = "rank:daily"
)

// PhotoViewKey 照片点击数的key
// view:photo:1 -> 100
// view:photo:2 -> 150
func PhotoViewKey(id uint) string {
	return fmt.Sprintf("view:photo:%s", strconv.Itoa(int(id)))
}
