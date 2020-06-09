package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"i-go/core/tbk"
	"i-go/demo/tbk/controller"
	"i-go/demo/tbk/server"
	"net/http"
)

func RegisterRouter(e *gin.Engine) {
	e.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "access denied!")
	})

	baseURL := e.Group("/api/v1")
	tbkGroup := baseURL.Group("/tbk")
	fmt.Println("conf:", tbk.TbkConf)
	tbkController := controller.NewTBK(server.NewTBK(tbk.TbkConf.AppKey, tbk.TbkConf.AppSecret, tbk.TbkConf.Router,
		tbk.TbkConf.Session, tbk.TbkConf.Timeout))
	// 淘宝客
	{
		tbkGroup.GET("", tbkController.FindURLByKeyWords)
	}
}
