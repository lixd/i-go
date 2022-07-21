package router

import (
	"i-go/core/db/mysqldb"
	"i-go/demo/user/controller"
	"i-go/demo/user/repository"
	"i-go/demo/user/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	e.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "access denied!")
	})

	baseURL := e.Group("/api/v1")
	user := baseURL.Group("/users")
	userController := controller.NewUser(server.NewUser(repository.NewUser(mysqldb.MySQL)))
	// 用户
	{
		user.POST("", userController.Insert)
		user.DELETE("", userController.Delete)
		user.PUT("", userController.Update)
		user.GET("/:id", userController.FindById)
		user.GET("", userController.Find)
	}
}
