package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"i-go/core/conf"
	"i-go/core/db/elasticsearch"
	"os"
)

var es *elastic.Client

func Init() {
	dir, _ := os.Getwd()
	fmt.Println("pwd: ", dir)
	err := conf.Load("conf/config.yml")
	if err != nil {
		panic(err)
	}
	elasticsearch.Init()
	es = elasticsearch.ESClient
}

func main() {
	Init()
	demoIndex := "twitter"
	exists, err := es.IndexExists(demoIndex).Do(context.Background())
	if err != nil {
		// Handle error
		logrus.Error(err)
	}
	if !exists {
		result, err := es.CreateIndex(demoIndex).Do(context.Background())
		if err != nil {
			// Handle error
			logrus.Error(err)
		}
		fmt.Println("CreateIndex ", result)
	} else {
		fmt.Println("index  exists")
	}
	es.IsRunning()
}
