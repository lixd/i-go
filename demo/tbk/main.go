package main

import (
	"github.com/gin-gonic/gin"
	"i-go/core/conf"
	"i-go/core/logger/ilogrus"
	"i-go/core/tbk"
	"i-go/demo/tbk/router"
)

func main() {
	Init()

	engine := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.RegisterRouter(engine)
	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}

// Init 初始化
/*
读取配置文件 初始化数据库连接
*/
func Init() {
	conf.Load("D:/lillusory/projects/i-go/conf/config.yml")
	_, err := tbk.ParseConf()
	if err != nil {
		panic(err)
	}
	ilogrus.InitLogger()
}
