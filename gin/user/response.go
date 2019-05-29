package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Gin 响应
type V4Controller struct {
}

func (v4Controller *V4Controller) StringHandler(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}

func (v4Controller *V4Controller) HTMLHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (v4Controller *V4Controller) JSONHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "{message:hello world}")
}
func (v4Controller *V4Controller) DATAHandler(c *gin.Context) {
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte("{message:hello world}"))
}
