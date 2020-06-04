package main

import (
	"github.com/gin-gonic/gin"
	"i-go/gin/vaptchademo/router"
)

func main() {
	g := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.RegisterRoutes(g)
	if err := g.Run(":8081"); err != nil {
		panic(err)
	}
}
