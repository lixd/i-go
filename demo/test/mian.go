package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	initConfig2()
}
func initConfig2() {
	viper.SetConfigFile("demo/conf/config.json") // 配置文件名 不加后缀
	// viper.AddConfigPath("./")     // 配置文件路径
	viper.SetConfigType("json") // 设置配置文件格式为json
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("conf init err %v \n", err)
	}
	appUrl := viper.GetString("mongo.appUrl")
	fmt.Printf("appurl=%v \n", appUrl)
}
