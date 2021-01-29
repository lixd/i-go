package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"i-go/gin/swagger/http/controller"
	// _ "i-go/gin/swagge/http/router/docs"
)

func RegisterHello(e *gin.Engine) {
	e.GET("/hello/:name", controller.Hello.Greeter)
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
