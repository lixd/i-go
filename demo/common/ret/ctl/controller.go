// Package ctl Controller层返回值
package ctl

import (
	"net/http"

	"i-go/demo/common/ret/srv"

	"github.com/gin-gonic/gin"
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
