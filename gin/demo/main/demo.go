package main

import (
	"github.com/gin-gonic/gin"
	"i-go/gin/demo/router"
)

func main() {
	engine := gin.Default()
	r := router.Router{}
	r.RegisterRouter(engine)
	engine.Run(":8080")
}
