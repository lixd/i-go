package router

import (
	"github.com/gin-gonic/gin"
	"i-go/core/db/mysqldb"
	"i-go/demo/order/controller"
	"i-go/demo/order/repository"
	"i-go/demo/order/server"
	"net/http"
)

func RegisterRouter(e *gin.Engine) {
	e.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "access denied!")
	})

	baseURL := e.Group("/api/v1")
	order := baseURL.Group("/order")
	orderController := controller.NewOrder(server.NewOrder(repository.NewOrder(mysqldb.MySQL)))
	// 订单
	{
		order.PUT("", orderController.Insert)
		order.DELETE("", orderController.Delete)
		order.POST("", orderController.Update)
		order.GET("", orderController.Find)
		order.GET("/list", orderController.FindList)
	}
}
