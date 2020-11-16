package time

import (
	"time"
)

const (
	Layout = "2006-01-02 15:04:05"
)

// DateTime 获取东八区时间戳(0点)
func DateTime() int64 {
	local := time.FixedZone("UTC", 8*3600)
	year, month, day := time.Now().In(local).Date()
	daytime := time.Date(year, month, day, 0, 0, 0, 0, local).Unix()
	timestamp := daytime
	return timestamp
}

// UnixTime 获取东八区时间戳
func UnixTime() int64 {
	loc := time.FixedZone("UTC", 8*3600)
	timestamp := time.Now().In(loc).Unix()
	return timestamp
}

func Format(unix int64) string {
	return time.Unix(unix, 0).Format(Layout)
}

func Parse(format string) (t time.Time) {
	t, _ = time.Parse(Layout, format)
	return
}
