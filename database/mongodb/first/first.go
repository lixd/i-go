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

// 5ed46396 1c7490 0ec8 482006
//

func main() {
	conf.Init("D:/lillusory/projects/i-go/conf/config.yml")
	mongodb.Init()

	upsert()
	//query()
	//aggregate()
}

func upsert() {
	var req = model.UserInfoReq{
		//Id:       "5ebd4fe4d8c4278a887c4539",
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
