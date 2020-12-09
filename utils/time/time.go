package time

import (
	"time"
)

const (
	Layout = "2006-01-02 15:04:05"
)

// GetDailyTime 获取东八区时间戳(0点)
func GetDailyTime(date time.Time) int64 {
	year, month, day := date.Date()
	daytime := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return daytime.Unix()
}

// GetUTC8 获取东八区时间戳
func GetUTC8() int64 {
	loc := time.FixedZone("UTC", 8*3600)
	timestamp := time.Now().In(loc).Unix()
	return timestamp
}

func Format(unix int64) string {
	return time.Unix(unix, 0).Format(Layout)
}
