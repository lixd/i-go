package search

import (
	"flag"
	"fmt"
	"testing"

	"i-go/core/conf"
	"i-go/core/db/elasticsearch"
	"i-go/core/logger"
)

func TestMain(m *testing.M) {
	var file string
	flag.StringVar(&file, "f", "conf/config_job.yaml", "the config file path")
	flag.Parse()

	if err := conf.Load(file); err != nil {
		panic(err)
	}
	logger.InitLogger()
	elasticsearch.Init()
	m.Run()
}

func TestSite_Search(t *testing.T) {
	total, data, err := SiteClient.Search("服装")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("总网址数:", total)
	fmt.Println(data)
}

func TestSite_SearchAll(t *testing.T) {
	data, err := SiteClient.SearchAll("服装")
	if err != nil {
		fmt.Println(err)
		return
	}
	var cnt int
	for v := range data {
		cnt++
		fmt.Println(v)
	}
	fmt.Println("总网址数:", cnt)
}
