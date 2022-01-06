package itime

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeUnix(t *testing.T) {
	// 1.时间戳和时区无关
	fmt.Println("当前时间戳-DEFAULT:", time.Now().Unix())
	fmt.Println("当前时间戳-UTC:", time.Now().UTC().Unix())
	fmt.Println("当前时间戳-Local:", time.Now().Local().Unix())
	// 2.相同时间戳在不同时区展示不一样
	fmt.Println("当前时间-DEFAULT:", time.Now().Format(time.RFC3339))
	fmt.Println("当前时间-UTC:", time.Now().UTC().Format(time.RFC3339))
	fmt.Println("当前时间-Local:", time.Now().Local().Format(time.RFC3339))
	// 3.当日零点时间-当前时区(东八区 UTC+8)的零点还是零时区(UTC+0)的零点
	// 比如 UTC+0 的零点到UTC+8就是八点了
	year, month, day := time.Now().Date()
	daytimeLocal := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	fmt.Println("零点时间戳-Local:", daytimeLocal.Unix())
	daytimeUTC := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	fmt.Println("零点时间戳-UTC0:", daytimeUTC.Unix())
}

func TestGetZeroTime(t *testing.T) {
	zeroTime := GetZeroTime(time.Now().Unix())
	fmt.Println("当日零点时间:", zeroTime.Unix())
}

func TestGetZeroTime2(t *testing.T) {
	zeroTime := getZeroTime2(time.Now().Unix())
	fmt.Println("当日零点时间:", zeroTime.Unix())
}

// 41.7 ns/op
func BenchmarkGetZeroTime(b *testing.B) {
	unix := time.Now().Unix()
	for i := 0; i < b.N; i++ {
		GetZeroTime(unix)
	}
}

// 240 ns/op
func BenchmarkGetZeroTime2(b *testing.B) {
	unix := time.Now().Unix()
	for i := 0; i < b.N; i++ {
		getZeroTime2(unix)
	}
}
