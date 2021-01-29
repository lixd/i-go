package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"i-go/core/conf"
	"i-go/core/db/mongodb"
	"i-go/database/mongodb/model"
	"i-go/database/mongodb/repository"
)

var tdb = mongodb.TestDB

func main() {
	conf.Load("D:/lillusory/projects/i-go/conf/config.yml")
	mongodb.Init()

	// upsert()
	incUpsert()
	// query()
	// aggregate()
}

func upsert() {
	var req = model.UserInfoReq{
		// Id:       "5ebd4fe4d8c4278a887c4539",
		UserName: "First",
		Password: "First",
		Age:      1323,
		Phone:    "13452340416",
		Hobby:    []string{"Reading", "Running", "Music"},
	}
	id, err := repository.UserInfo.Upsert(&req)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "mongodb 插入数据失败"}).Error(err)
	}
	fmt.Println(id)
}

func incUpsert() {
	err := repository.UserInfo.IncUpsert()
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "mongodb 插入数据失败"}).Error(err)
	}
}

func query() {
	infos, err := repository.UserInfo.QueryByHobby("Reading")
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "mongodb 查询失败"}).Error(err)
	}
	fmt.Println(infos)
}

func aggregate() {
	infos, err := repository.UserInfo.QueryCount("First")
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "mongodb 查询失败"}).Error(err)
	}
	fmt.Println(infos)
}
