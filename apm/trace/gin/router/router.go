package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"time"
)

func Register(e *gin.Engine) {
	e.GET("/ping", Ping)
}

func Ping(c *gin.Context) {
	psc, _ := c.Get("ctx")
	ctx := psc.(context.Context)
	doPing1(ctx)
	doPing2(ctx)
	c.JSON(200, gin.H{"message": "pong"})
}
func doPing1(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "doPing1")
	defer span.Finish()
	time.Sleep(time.Second)
	fmt.Println("pong")
}
func doPing2(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "doPing2")
	defer span.Finish()
	time.Sleep(time.Second)
	fmt.Println("pong")
}
