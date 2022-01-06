package itime

import (
	"time"
)

/*
时间戳 在地球的每一个角落都是相同的
时区 时间的表示方式

 1.时间戳和时区无关
 2.相同时间戳在不同时区以不同的形式展现
 3.时间字符串转成时间戳时需要指定时区。因为Unix是没带时区的，但是时间字符串中带了，如果解析时不指定则会丢失时区信息。
参考：https://www.jianshu.com/p/bf47458a0423
*/

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
