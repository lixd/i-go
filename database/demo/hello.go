package main

import (
	"fmt"
	"i-go/core/conf"
	"i-go/core/db/mysqldb"
	"os"
)

func main() {
	fmt.Println(os.Getwd())
	//conf.Load("/conf/config.yml")
	conf.Load("D:/lillusory/projects/i-go/conf/config.yml")
	mysqldb.Init()
}
