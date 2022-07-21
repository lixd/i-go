package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"i-go/apm/trace/gin/middleware"
	"i-go/apm/trace/gin/router"

	"github.com/gin-gonic/gin"
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

	// 等待中断信号以优雅地关闭服务器
	signalChan := make(chan os.Signal, 1)
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
