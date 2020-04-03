package utils

import "time"

// 获取东八区零时的时间戳
func GetDailyUnix(date time.Time) (daytime int64) {
	local := time.FixedZone("UTC", 8*3600)
	time.Unix(date.Unix(), 0).In(local)
	year, month, day := date.Date()
	daytime = time.Date(year, month, day, 0, 0, 0, 0, local).Unix()
	return
}
