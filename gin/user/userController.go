package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type V1 interface {
	Index(c *gin.Context)
	Ping(c *gin.Context)
	User(c *gin.Context)
	Redirect(c *gin.Context)
}

type v1Controller struct {
}

var V1C = &v1Controller{}

func (v1 *v1Controller) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// hello world
func (v1 *v1Controller) Ping(c *gin.Context) {
	// gin.H 是 map[string]interface{} 的一种快捷方式
	c.JSON(200, gin.H{"message": "pong"})
}

// AsciiJSON
func (v1 *v1Controller) User(c *gin.Context) {
	User1 := User{"illusory", 23, "CQ"}
	c.AsciiJSON(http.StatusOK, User1)
}

func (v1 *v1Controller) Redirect(c *gin.Context) {
	log.Print("HTTP重定向。。")
	c.Redirect(http.StatusMovedPermanently, "https://www.lixueduan.com")
}
