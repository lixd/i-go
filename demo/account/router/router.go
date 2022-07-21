package router

import (
	"i-go/core/db/mysqldb"
	"i-go/demo/account/controller"
	"i-go/demo/account/repository"
	"i-go/demo/account/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	e.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "access denied!")
	})

	baseURL := e.Group("/api/v1")
	account := baseURL.Group("/accounts")
	accountController := controller.NewAccount(server.NewAccount(repository.NewAccount(mysqldb.MySQL)))
	// 订单
	{
		account.POST("", accountController.Insert)
		account.DELETE("", accountController.DeleteByUserId)
		account.PUT("", accountController.Update)
		account.GET("/:userId", accountController.FindByUserId)
		account.GET("", accountController.FindList)
	}
}
