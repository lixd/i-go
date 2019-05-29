package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"i-go/src/gin/routers"
	"log"
	"net/http"
	"time"
)

func main() {
	// Default方法创建一个路由handler。
	router := gin.Default()
	routers.RegisterRoutes(router)
	// 修改模式
	// gin.SetMode(gin.ReleaseMode)
	// 记录日志到文件
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	// 	log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	// }
	router.Run(":8080")

	// 启动两个服务
	// RunTwoServer()
}

func RunTwoServer() {
	var (
		g errgroup.Group
	)
	server01 := &http.Server{
		Addr:              ":8081",
		Handler:           router01(),
		ReadTimeout:       time.Second * 5,
		ReadHeaderTimeout: time.Second * 10}
	server02 := &http.Server{
		Addr:              ":8082",
		Handler:           router02(),
		ReadTimeout:       time.Second * 5,
		ReadHeaderTimeout: time.Second * 10}
	g.Go(func() error {
		return server01.ListenAndServe()
	})
	g.Go(func() error {
		return server02.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 01",
			},
		)
	})
	return e
}
func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 02",
			},
		)
	})
	return e
}
