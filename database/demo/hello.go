package main

import (
	"fmt"
	"i-go/core/conf"
	"i-go/core/db/mysqldb"
	"os"
)

func main() {
	fmt.Println(os.Getwd())
	//conf.Init("/conf/config.yml")
	conf.Init("D:/lillusory/projects/i-go/conf/config.yml")
	mysqldb.Init()
}
