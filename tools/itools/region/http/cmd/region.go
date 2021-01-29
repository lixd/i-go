package main

import (
	"github.com/gin-gonic/gin"
	"i-go/tools/itools/region/http/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()
	engine.Use(gin.Recovery())

	router.RegisterRegion(engine)
	if err := engine.Run(":8081"); err != nil {
		panic(err)
	}
}
