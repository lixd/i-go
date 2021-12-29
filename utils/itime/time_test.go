package itime

import (
	"fmt"
	"testing"
	"time"
)

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
