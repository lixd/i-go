package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"i-go/core/conf"
	"i-go/core/db/mongodb"
	"i-go/database/mongodb/model"
	"i-go/database/mongodb/repository"
)

var tdb = mongodb.TestDB

func main() {
	conf.Load("D:/lillusory/projects/i-go/conf/config.yml")
	mongodb.Init()
	// arrayUpsert()
	arrayUpsert2()

	// upsert()
	// incUpsert()
	// query()
	// aggregate()
}

func arrayUpsert() {
	filter := bson.M{"UserName": "Reselect"}
	update := bson.M{
		"$set": bson.M{
			"Reselect":   []model.Reselect{{Node: "hk", Status: 1}},
			"UpdateTime": time.Now().Unix(),
		},
		"$setOnInsert": bson.M{"CreateTime": time.Now().Unix()},
	}
	opts := options.Update().SetUpsert(true)
	_, err := repository.UserInfo.GetColl().UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		logrus.Error(err)
	}
}

func arrayUpsert2() {
	// filter := bson.M{"UserName": "Reselect", "Reselect": bson.M{"$elemMatch": bson.M{"Node": "hk"}}}
	filter := bson.M{"UserName": "Reselect", "Reselect.Node": "hk"}
	update := bson.M{
		"$set": bson.M{"Reselect.$.Status": 1}}
	opts := options.Update().SetUpsert(true)
	_, err := repository.UserInfo.GetColl().UpdateMany(context.Background(), filter, update, opts)
	if err != nil {
		logrus.Error(err)
	}
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
