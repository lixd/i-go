package ratelimit

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/aegis/ratelimit"
)

func Test_limitDemo(t *testing.T) {
	limitDemo()
}

func TestRateLimit(t *testing.T) {
	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			fmt.Printf("限流状态：%+v\n", limiter.Stat())
		}
	}()
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.GET("/limit", limit(), logic)
	r.GET("/nolimit", logic)
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

func limit() gin.HandlerFunc {
	return func(c *gin.Context) {
		doneFunc, err := limiter.Allow()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, "request-limit")
			return
		}
		defer doneFunc(ratelimit.DoneInfo{})
		c.Next()
	}
}

func logic(c *gin.Context) {
	// 差不过50ms
	for j := 0; j < 1000_0000; j++ {
	}
	_, _ = c.Writer.Write([]byte("Hello,World"))
}

/*
压测结果
 ./wrk -t 100 -c 1500 -d 60s --timeout 5s http://localhost:8080/limit
 ./wrk -t 100 -c 1500 -d 300s http://localhost:8080/nolimit
不加limit的情况
  18288 requests in 5.00m, 2.23MB read
  Socket errors: connect 579, read 0, write 0, timeout 18228
基本全部超时，只有开始的几十个成功了
带limit的情况
  29706 requests in 5.00m, 3.86MB read
  Socket errors: connect 579, read 0, write 0, timeout 19919
还是超时很多，不过能正常处理一部分
*/
