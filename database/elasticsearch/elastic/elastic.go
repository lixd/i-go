package main

import (
	"i-go/core/conf"
	"i-go/core/db/elasticsearch"
)

var es = elasticsearch.ESClient

func Init() {
	conf.Init("D:/lillusory/projects/i-go/conf/config.yml")
	elasticsearch.Init()
}

func main() {
	Init()
}
