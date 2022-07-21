package main

import (
	"io"
	"log"
	"os"

	router "i-go/gin/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Default方法创建一个路由handler。
	e := gin.Default()
	router.RegisterRoutes(e)
	// 修改模式
	gin.SetMode(gin.DebugMode)
	// 记录日志到文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	if err := e.Run(":8080"); err != nil {
		panic(err)
	}
}
