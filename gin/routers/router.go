package router

import (
	"net/http"

	"i-go/gin/user"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// 加载静态资源
	router.LoadHTMLFiles("D:/lillusory/projects/i-go/gin/templates/advert.tmpl")
	// router.LoadHTMLGlob("../templates/*")
	// 设定请求url不存在的返回值
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "page not found"})
	})
	// 重定向
	router.GET("/redirect", user.V1C.Redirect)
	// 路由重定向
	router.GET("/routeredirect", func(c *gin.Context) {
		c.Request.URL.Path = "/redirect"
		router.HandleContext(c)
	})

	// route 分组
	v1 := router.Group("/v1")
	// 使用中间件
	// v1.Use(middleware.AuthMiddleWare())
	{
		v1.GET("/", user.V1C.Index)
		v1.GET("/ping", user.V1C.Ping)
		v1.GET("/user", user.V1C.User)
	}

}
