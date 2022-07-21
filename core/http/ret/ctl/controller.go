// Controller层返回值
package ctl

import (
	"i-go/core/http/ret/svc"

	"github.com/gin-gonic/gin"

	"net/http"
)

func BadRequest(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusOK, svc.BadRequest(msg...))
}

func Unauthorized(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, svc.Unauthorized())
}

func Forbidden(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, svc.Forbidden())
}

func Limit(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, svc.Limit())
}

func Success(c *gin.Context, result *svc.Result) {
	c.JSON(http.StatusOK, result)
}
