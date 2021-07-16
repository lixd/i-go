package ratelimit

import (
	"fmt"
	"testing"
	"time"

	"go.uber.org/ratelimit"
)

func TestLeakyBucket(t *testing.T) {
	// 限制每秒 100 次请求 即 10ms 一次
	rl := ratelimit.New(100) // per second

	prev := time.Now()
	for i := 0; i < 1000; i++ {
		now := rl.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}

/*
// uber ratelimit 部分代码
func (t *limiter) Take() {
	// 1.首先获取当前时间
	now := t.clock.Now()
	// 2.获取需要 sleep 的时间 用每次请求的时间减去上次请求到现在经过的时间
	// 其中 perRequest 计算为 perRequest=time.Second / time.Duration(rate),
	t.sleepFor += t.perRequest - now.Sub(t.last)
	// 3. 最后就 sleep 一段时间
	if t.sleepFor > 0 {
		t.clock.Sleep(t.sleepFor)
	}
}
*/
