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
	order := baseURL.Group("/orders")
	orderController := controller.NewOrder(server.NewOrder(repository.NewOrder(mysqldb.MySQL)))
	// 订单
	{
		order.POST("", orderController.Insert)
		order.DELETE("", orderController.Delete)
		order.PUT("", orderController.Update)
		order.GET("/:id", orderController.FindById)
		order.GET("", orderController.Find)
	}
}
