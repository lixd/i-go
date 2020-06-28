package router

import (
	"github.com/gin-gonic/gin"
	"i-go/gin/vaptchademo/logic"
	"net/http"
	"os"
	"time"
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
	// JSONP lang 接口
	g.GET("/api/v1/lang", func(context *gin.Context) {
		time.Sleep(time.Second * 10)
		data := map[string]string{
			"lang": "zh-TW",
		}
		context.JSONP(http.StatusOK, data)
	})
}
