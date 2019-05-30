package router

import (
	"github.com/gin-gonic/gin"
	"i-go/gin/demo/controller"
)

type Router struct {
}

func (r *Router) RegisterRouter(e *gin.Engine) {
	demo := e.Group("/demo")
	accountController := &controller.AccountController{}
	{
		demo.GET("/login", accountController.LoginHandler)
	}
}
