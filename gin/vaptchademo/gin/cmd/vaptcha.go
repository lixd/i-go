package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"i-go/gin/vaptchademo/gin/logic"
)

// http://localhost:8080/vaptcha/demo/click.html
func main() {
	g := gin.Default()
	v := g.Group("/vaptcha")
	{
		v.POST("/login", logic.Login)
		v.GET("/offline", logic.Offline)
		// 加载静态资源 注意调整路径
		v.StaticFS("/demo", http.Dir("../../assets"))
	}
	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Run error:%v", err)
	}
}
