package ratelimit

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/aegis/ratelimit"
	"github.com/go-kratos/aegis/ratelimit/bbr"
	"google.golang.org/grpc"
)

func main() {
	limiter := bbr.NewLimiter()
	doneFunc, err := limiter.Allow()
	if err != nil { // 返回错误则说明需要进行限流
		return
	}
	defer doneFunc(ratelimit.DoneInfo{}) // 最好在 defer 中执行 doneFunc,保证panic之后也不会导致数据异常
	// 	your logic here
}

func limitDemo() {
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
			time.Sleep(time.Millisecond * 100)
		} else {
			go func() {
				defer doneFunc(ratelimit.DoneInfo{})
				for j := 0; j < 100_0000; j++ {

				}
			}()
		}
	}
	time.Sleep(time.Second)
}

var limiter = bbr.NewLimiter()

// RateLimit is a server interceptor that rate limits requests.
func RateLimit(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	doneFunc, err := limiter.Allow()
	if err != nil {
		return nil, err
	}
	defer doneFunc(ratelimit.DoneInfo{})
	return handler(ctx, req)
}

// RateLimit2 is a gin middleware that rate limits requests.
func RateLimit2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		doneFunc, err := limiter.Allow()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, "request-limit")
			return
		}
		defer doneFunc(ratelimit.DoneInfo{})
		ctx.Next()
	}
}
