package router

import (
	"github.com/gin-gonic/gin"
	"i-go/gin/vaptchademo/logic"
	"net/http"
	"os"
)

func RegisterRoutes(g *gin.Engine) {
	// 加载静态资源
	dir, _ := os.Getwd()
	println("pwd: ", dir)
	// 注意调整路径
	g.StaticFS("/static", http.Dir("./gin/vaptchademo/static"))
	//g.StaticFS("/static", http.Dir("./static"))
	vaptcha := g.Group("/vaptcha")
	vaptcha.GET("/offline", logic.VaptchaDemo.Offline)
	vaptcha.POST("/login", logic.VaptchaDemo.Login)
}
