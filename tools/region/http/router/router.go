package router

import (
	"i-go/tools/region/http/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRegion(e *gin.Engine) {
	tools := e.Group("/api/v1/tools/")
	tools.GET("/ip2region", controller.Tools.Ip2Region)
	tools.GET("/ip2latlong", controller.Tools.Ip2LatLong)
}
