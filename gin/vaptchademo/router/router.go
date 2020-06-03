package router

import (
	"github.com/gin-gonic/gin"
	"i-go/gin/vaptchademo/logic"
	"net/http"
	"os"
)

func RegisterRoutes(g *gin.Engine) {
	dir, _ := os.Getwd()
	println("pwd: ", dir)
	// 加载静态资源 注意调整路径
	g.StaticFS("/static", http.Dir("./gin/vaptchademo/static"))
	//g.StaticFS("/static", http.Dir("./static"))
	vaptcha := g.Group("/vaptcha")
	vaptcha.GET("/offline", logic.VaptchaDemo.Offline) // 离线验证
	vaptcha.POST("/login", logic.VaptchaDemo.Login)    // 登录+二次验证
}
