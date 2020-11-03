// Controller层返回值
package ctl

import (
	"github.com/gin-gonic/gin"
	"i-go/demo/common/ret/srv"
	"net/http"
)

func BadRequest(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, srv.BadRequest())
}

func Unauthorized(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, srv.Unauthorized())
}

func Forbidden(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, srv.Forbidden())
}

func Success(c *gin.Context, result *srv.Result) {
	c.JSON(http.StatusOK, result)
}
