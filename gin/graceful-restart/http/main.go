package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// Go1.18后提供的http.Server.Shutdown方法实现了优雅关机

func main() {
	// mux demo
	// r := mux.NewRouter()
	// r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	// 	time.Sleep(5 * time.Second)
	// 	_, _ = writer.Write([]byte("mux ok"))
	// }).Methods("GET")

	// gin demo
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.JSON(http.StatusOK, "gin ok")
	})
	// 注：这里需要创建一个http.Server对象，然后调用server.ListenAndServe方法，而不是直接调用http.ListenAndServe方法.
	// 因为graceful shutdown方法时由http.Server对象实现的
	// http.ListenAndServe内部也是创建了一个http.Server对象，只是外部无法获取
	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server listen err:%s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, channel := context.WithTimeout(context.Background(), 5*time.Second)
	defer channel()
	// 优化关闭服务--会在请求处理完成后再关闭服务，而不是立即关闭
	// 通过ctx添加一个5秒钟超时限制
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown error")
	}
	log.Println("server exiting...")
}
