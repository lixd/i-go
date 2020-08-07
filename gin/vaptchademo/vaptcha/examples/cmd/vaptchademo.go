package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"i-go/gin/vaptchademo/router"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	g := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.RegisterRoutes(g)
	pprof.Register(g, "debug/pprof")
	go func() {
		http.ListenAndServe(":8000", nil)
	}()
	if err := g.Run(":8080"); err != nil {
		panic(err)
	}
}
