package main

import (
	"github.com/gin-gonic/gin"
	"i-go/core/conf"
	"i-go/tools/region/core"
	"i-go/tools/region/http/router"
)

func main() {
	err := conf.Load("conf/config.yml")
	if err != nil {
		panic(err)
	}
	core.Init()

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	router.RegisterRegion(engine)
	if err := engine.Run(":8081"); err != nil {
		panic(err)
	}
}
