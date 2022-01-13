package ratelimit

import (
	"fmt"
	"log"
	"time"

	"github.com/go-kratos/aegis/ratelimit"
	"github.com/go-kratos/aegis/ratelimit/bbr"
)

func circuitBreaker() {
	limiter := bbr.NewLimiter()
	go func() {
		ticker := time.NewTicker(time.Millisecond * 100)
		for range ticker.C {
			fmt.Printf("限流状态：%+v\n", limiter.Stat())
		}
	}()
	for i := 0; i < 100_0000; i++ {
		// allow 方法返回一个 func,请求处理完成后需要执行该 func 以告知 limiter，使其同步数据
		doneFunc, err := limiter.Allow()
		if err != nil {
			log.Println(err)
		} else {
			go func() {
				for j := 0; j < 100_0000; j++ {

				}
				doneFunc(ratelimit.DoneInfo{})
			}()
		}
	}
	time.Sleep(time.Second)
}
