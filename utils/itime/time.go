package itime

import (
	"time"
)

/*
时间戳 在地球的每一个角落都是相同的
时区 时间的表示方式
*/
const (
	layout = "2006-01-02 15:04:05"
)

// GetZeroTime 获取指定时间戳对应日期的零点的时间
func GetZeroTime(unix int64) time.Time {
	year, month, day := time.Unix(unix, 0).Date()
	daytime := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return daytime
}

// getZeroTime2 获取当天零点的时间比较慢，不推荐使用
func getZeroTime2(unix int64) time.Time {
	timeStr := time.Unix(unix, 0).Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return t
}

// Format 格式化时间
func Format(unix int64) string {
	return time.Unix(unix, 0).Format(layout)
}
