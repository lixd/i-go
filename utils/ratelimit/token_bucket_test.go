package ratelimit

import (
	"context"
	"fmt"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestTokenBucketAllow(t *testing.T) {
	// 1st: 每秒新增 token 数
	// 2ed: 桶中令牌最大数
	// 每次请求消耗一枚 token
	limiter := rate.NewLimiter(1, 2)
	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			fmt.Println("通过")
		} else {
			fmt.Println("wait...")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func TestTokenBucketWait(t *testing.T) {
	// 1st: 每秒新增 token 数
	// 2ed: 桶中令牌最大数
	// 每次请求消耗一枚 token
	limiter := rate.NewLimiter(1, 1)
	for i := 0; i < 10; i++ {
		err := limiter.Wait(context.Background())
		if err != nil {
			fmt.Println("err:", err)
			time.Sleep(time.Millisecond * 500)
			continue
		}
	}
}

/*
//rate包部分代码
func reserveN(n int){
	// 1.计算当前的 token 数
	// 也是按时间计算 当前时间和上次请求的时间这中间的时候 又增加了多少 token 数。
	now, last, tokens := lim.advance(now)

	// 2.然后减去 当前需要消耗的 token 数
	tokens -= float64(n)

	// 3.然后如果减完是负数则计算一下多消耗的令牌需要多久才能生成出来
	// 另外一个 Wait 方法是阻塞的 这里算出来应该就是阻塞时间
	var waitDuration time.Duration
	if tokens < 0 {
		waitDuration = lim.limit.durationFromTokens(-tokens)
	}

	// 4.这里的 ok 就是最终返回的值
	// 如果 需要消耗的令牌数直接都大于了桶容量(burst) 那肯定是 false
	// 如果 需要等待的时间 waitDuration 大于 指定时间(maxFutureReserve) 也是 false
	ok := n <= lim.burst && waitDuration <= maxFutureReserve
}*/
