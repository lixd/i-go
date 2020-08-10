package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"i-go/gin/vaptchademo/gin/logic"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	g := gin.Default()
	v := g.Group("/vaptcha")
	v.POST("/login", logic.Login)
	v.GET("/offline", logic.Offline)
	// 加载静态资源 注意调整路径
	// g.StaticFS("/static", http.Dir(buildPath()))
	g.StaticFS("/static", http.Dir("E:\\lillusory\\Projects\\i-go\\gin\\vaptchademo\\gin\\static"))
	if err := g.Run(":8080"); err != nil {
		panic(err)
	}
}

// buildPath build path to resource
func buildPath() (path string) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir, _ = filepath.Split(dir)
	path = filepath.Join(dir, "static")
	fmt.Println("path:", path)
	return path
}
