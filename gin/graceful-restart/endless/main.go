package main

import (
	"log"
	"net/http"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

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
		c.String(http.StatusOK, "gin ok4")
	})
	// 默认endless服务会监听下列信号：
	// syscall.SIGHUP，syscall.SIGUSR1，syscall.SIGUSR2，syscall.SIGINT，syscall.SIGTERM和syscall.SIGTSTP
	// 接收到 SIGHUP 信号将触发`fork/restart` 实现优雅重启（kill -1 pid会发送SIGHUP信号）
	// 接收到 syscall.SIGINT或syscall.SIGTERM 信号将触发优雅关机
	// 接收到 SIGUSR2 信号将触发HammerTime
	// SIGUSR1 和 SIGTSTP 被用来触发一些用户自定义的hook函数
	server := endless.NewServer(":8080", r)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("listen: %s\n", err)
	}

	log.Println("Server exiting")

}
