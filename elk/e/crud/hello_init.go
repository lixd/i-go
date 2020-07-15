package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"i-go/core/conf"
	"i-go/core/db/elasticsearch"
	"i-go/utils"
)

const (
	HelloIndex   = "hello"
	HelloIndexV1 = "hello_v1"
	HelloMapping = `
					{
						"settings": {
						 "number_of_shards": 3,
						 "number_of_replicas": 1
					  },
					  "mappings" : {
						"properties" : {
						  "id" : {
							"type" : "long"
						  },
						  "name" : {
							"type" : "text",
							"fields" : {
							  "keyword" : {
								"type" : "keyword",
								"ignore_above" : 256
							  }
							}
						  }
						}
					  }
					}
`
)

func Init() {
	// err := conf.Init("D:/lillusory/projects/i-go/conf/config.yml")
	err := conf.Init("./conf/config.yml")
	if err != nil {
		panic(err)
	}
	elasticsearch.Init()
}

func main() {
	Init()
	h := hello{
		Index: HelloIndex,
		Type:  "_doc",
		Cli:   elasticsearch.ESClient,
	}
	h.CreateIndex()
	h.AddAlias("", HelloIndexV1)
}

func (s *hello) CreateIndex() {
	isExists, err := s.Cli.IndexExists(HelloIndex).Do(context.Background())
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller()}).Error(err)
		return
	}
	if !isExists {
		logrus.Println("index does not exists,will create")
		s.createIndex()
	} else {
		logrus.Info("index already exists")
	}
}

func (s *hello) createIndex() {
	_, err := s.Cli.CreateIndex(HelloIndex).
		BodyJson(HelloMapping).
		Do(context.Background())
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller()}).Error(err)
		return
	}
	logrus.Info("create index successful")
}

// alias remove index alias from old and add to new
func (s *hello) AddAlias(old, new string) {
	alias := s.Cli.Alias()
	if old != "" {
		alias = alias.Remove(old, HelloIndex)
	}
	_, err := alias.
		Add(new, HelloIndex).
		Do(context.Background())
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller()}).Error(err)
		return
	}
	logrus.Info("add index alias successful")
}
