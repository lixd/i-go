package router

import (
	"github.com/gin-gonic/gin"
	"i-go/tools/itools/region/http/controller"
)

func RegisterRegion(e *gin.Engine) {
	tools := e.Group("/api/v1/tools/")
	tools.GET("/ip2region", controller.Tools.Ip2Region)
}
