package router

import (
	"github.com/gin-gonic/gin"
	"i-go/core/db/mysqldb"
	"i-go/demo/account/controller"
	"i-go/demo/account/repository"
	"i-go/demo/account/server"
	"net/http"
)

func RegisterRouter(e *gin.Engine) {
	e.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "access denied!")
	})

	baseURL := e.Group("/api/v1")
	account := baseURL.Group("/account")
	accountController := controller.NewAccount(server.NewAccount(repository.NewAccount(mysqldb.MySQL)))
	// 订单
	{
		account.PUT("", accountController.Insert)
		account.DELETE("", accountController.Delete)
		account.POST("", accountController.Update)
		account.GET("", accountController.Find)
		account.GET("/list", accountController.FindList)
	}
}
