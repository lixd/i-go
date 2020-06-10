package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func TestNewRequestLimitService(t *testing.T) {
	rater := NewRateLimiter(time.Second*10, 10)
	for i := 0; i < 10; i++ {
		if rater.AllowN(2) {
			fmt.Printf("通过 当前请求数:%v \n", rater.ReqCount())
		} else {
			fmt.Println("wait...")
			time.Sleep(time.Millisecond * 500)
		}
	}
}
