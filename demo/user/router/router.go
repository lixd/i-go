package router

import (
	"github.com/gin-gonic/gin"
	"i-go/core/db/mysqldb"
	"i-go/demo/user/controller"
	"i-go/demo/user/repository"
	"i-go/demo/user/server"
	"net/http"
)

func RegisterRouter(e *gin.Engine) {
	e.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "access denied!")
	})

	baseURL := e.Group("/api/v1")
	user := baseURL.Group("/user")
	userController := controller.NewUser(server.NewUser(repository.NewUser(mysqldb.MySQL)))
	baseURL.POST("/post", userController.Test)
	// 用户
	{
		user.PUT("", userController.Insert)
		user.DELETE("", userController.Delete)
		user.POST("", userController.Update)
		user.GET("", userController.Find)
		user.GET("/list", userController.FindList)
	}
}
