package router

import (
	"github.com/gin-gonic/gin"
	"i-go/demo/controller"
	"i-go/demo/repository"
	"i-go/demo/server"
)

type Router struct {
}

func (r *Router) RegisterRouter(e *gin.Engine) {
	accountDAO := repository.AccountDAO{}
	dao := repository.DAO(&accountDAO)
	accountServer := server.AccountServer{Dao: &dao}
	server1 := server.Server(&accountServer)
	accountController := controller.AccountController{Server: &server1}

	demo := e.Group("/demo")
	{
		demo.GET("/login", accountController.LoginHandler)
		demo.POST("/register", accountController.RegisterHandler)
	}
}
