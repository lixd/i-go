package ratelimit

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"go.uber.org/ratelimit"
)

// go.uber.org/ratelimit 漏桶算法使用
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
// uber ratelimit 漏桶算法-部分核心代码
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

func Test_leakyBucket_allow(t *testing.T) {
	l := newLeakyBucket(4, 1)
	for i := 0; i < 10; i++ {
		if l.allow() {
			fmt.Println("通过")
		} else {
			fmt.Println("wait...")
		}
		time.Sleep(time.Millisecond * 300)
	}

	fmt.Println(strings.Repeat("~", 20))
	time.Sleep(time.Second * 2) // sleep 两秒,两秒后漏桶中最少已经空出两个位置了，后续2个请求可以立马执行
	for i := 0; i < 2; i++ {
		if l.allow() {
			fmt.Println("通过")
		} else {
			fmt.Println("wait...")
		}
	}
}
