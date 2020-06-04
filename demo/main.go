package main

import (
	"github.com/gin-gonic/gin"
	"i-go/core/conf"
	"i-go/core/db/mysqldb"
	"i-go/core/logger/izap"
	amodel "i-go/demo/account/model"
	arouter "i-go/demo/account/router"
	omodel "i-go/demo/order/model"
	orouter "i-go/demo/order/router"
	umodel "i-go/demo/user/model"
	urouter "i-go/demo/user/router"
)

func main() {
	Init()

	engine := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	urouter.RegisterRouter(engine)
	arouter.RegisterRouter(engine)
	orouter.RegisterRouter(engine)
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
	mysqldb.MySQL.AutoMigrate(&umodel.User{})
	mysqldb.MySQL.AutoMigrate(&amodel.Account{})
	mysqldb.MySQL.AutoMigrate(&omodel.Order{})

	//ilogrus.InitLogger()
	izap.InitLogger()
}
