package main

import (
	"i-go/core/conf"
)

func main() {
	path1 := "conf/api/elasticsearch.yaml"
	path2 := "conf/api/mongodb.yaml"
	conf.Loads([]string{path1, path2})
	select {}
}
