package user

import (
	"github.com/gin-gonic/gin"
	"log"
)

type V5Controller struct {
	int
}

func (v5Controller *V5Controller) Get(c *gin.Context) {
	c.Set("name", "admin")
	// Context 内部一个存放key-value的 map Keys map[string]interface{}
	// Keys is a key/value pair exclusively for the context of each request.
	value, exists := c.Get("name")
	log.Printf("c.Get() value:=%v exists:=%v", value, exists)
}
