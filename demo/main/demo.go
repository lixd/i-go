package main

import (
	"github.com/gin-gonic/gin"
	"i-go/demo/router"
	_ "net/http/pprof"
)

func main() {
	engine := gin.Default()
	r := router.Router{}
	r.RegisterRouter(engine)

	engine.Run(":8080")
}
