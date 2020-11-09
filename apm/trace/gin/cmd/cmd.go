package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"i-go/apm/trace/gin/middleware"
	"i-go/apm/trace/gin/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	e := gin.New()
	e.Use(middleware.Jaeger())
	router.Register(e)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      e,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
