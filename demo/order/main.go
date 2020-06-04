package main

import (
	"github.com/gin-gonic/gin"
	"i-go/core/conf"
	"i-go/core/db/mysqldb"
	"i-go/core/logger/ilogrus"
	"i-go/demo/order/model"
	"i-go/demo/order/router"
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
	conf.Init("D:/lillusory/projects/i-go/conf/config.yml")
	mysqldb.Init()
	mysqldb.MySQL.AutoMigrate(&model.Order{})

	ilogrus.InitLogger()
}
