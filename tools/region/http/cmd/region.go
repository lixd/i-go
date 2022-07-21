package main

import (
	"fmt"

	"i-go/core/conf"
	"i-go/tools/region/core"
	"i-go/tools/region/http/router"

	"github.com/gin-gonic/gin"
)

func main() {
	err := conf.Load("conf/config_ip.yaml")
	if err != nil {
		panic(err)
	}
	core.InitRegion()
	core.InitLatLong()

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	router.RegisterRegion(engine)
	fmt.Println("HTTP Server Is Running")
	if err := engine.Run(":8081"); err != nil {
		panic(err)
	}
}
