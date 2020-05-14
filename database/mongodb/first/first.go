package main

import (
	"github.com/sirupsen/logrus"
	"i-go/core/db/mongodb"
	"i-go/database/mongodb/model"
	"i-go/database/mongodb/repository"
)

var tdb = mongodb.TestDB

func main() {
	insert()
}

func insert() {
	var req = model.UserInfoReq{
		ID:       "5ebd4fe4d8c4278a887c4539",
		UserName: "First",
		Password: "First",
		Age:      133,
		Phone:    "13452340416",
	}
	err := repository.UserInfo.Upsert(&req)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "mongodb 插入数据失败"}).Error(err)
	}
}
