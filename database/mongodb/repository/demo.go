package repository

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"i-go/core/db/mongodb"
	"i-go/database/mongodb/model"
	"time"
)

type userInfo struct {
	coll *mongo.Collection
}

var UserInfo = userInfo{}

func (ui *userInfo) GetColl() *mongo.Collection {
	if ui.coll == nil {
		ui.coll = mongodb.GetTestCollection(new(model.UserInfo))
	}
	return ui.coll
}

type Template struct {
	Hour int `bson:"Hour"`
}

func (ui *userInfo) IncUpsert() error {
	list := make([]Template, 0, 24)
	for i := 0; i < 24; i++ {
		list = append(list, Template{Hour: i})
	}
	filter := bson.M{"UserName": "17x"}
	update := bson.M{
		"$addToSet": bson.M{"Hours24": bson.M{"$each": list}},
		// "$inc":  bson.M{"Age": 1, "Hours24.0.request": 1},
		"$set": bson.M{"Phone": 12345},
	}
	opts := options.Update().SetUpsert(true)
	_, err := ui.GetColl().UpdateOne(context.Background(), filter, update, opts)
	return err
}

func (ui *userInfo) Upsert(req *model.UserInfoReq) (string, error) {
	var (
		filter   bson.M
		objectID primitive.ObjectID
		err      error
	)

	if len(req.ID) != 0 {
		objectID, err = primitive.ObjectIDFromHex(req.ID)
		if err != nil {
			logrus.WithFields(logrus.Fields{"Scenes": "upsert失败"}).Error(err)
		}
	} else {
		objectID = primitive.NewObjectID()
	}
	filter = bson.M{"_id": objectID}

	update := bson.M{
		"$setOnInsert": bson.M{"CreateTime": time.Now().Unix()},
		"$set": bson.M{
			"UserName":   req.UserName,
			"Password":   req.Password,
			"Age":        req.Age,
			"Phone":      req.Phone,
			"Hobby":      req.Hobby,
			"UpdateTime": time.Now().Unix(),
		},
	}
	opts := options.Update().SetUpsert(true)
	_, err = ui.GetColl().UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		logrus.Error(err)
	}
	return objectID.Hex(), nil
}

func (ui *userInfo) QueryByHobby(hobby string) ([]model.UserInfo, error) {
	var (
		list = make([]model.UserInfo, 0)
		item = model.UserInfo{}
	)
	filter := bson.M{"Hobby": bson.M{"$elemMatch": bson.M{"$eq": hobby}}}
	cursor, err := ui.GetColl().Find(context.Background(), filter)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		if err := cursor.Decode(&item); err != nil {
			logrus.WithFields(logrus.Fields{"Scenes": "decode error"}).Error(err)
			continue
		}
		list = append(list, item)
	}
	return list, nil
}

// 每小时从攻击日志中统计出{区域、类型、活跃攻击ip}数量
type QueryCount struct {
	Count int64 `bson:"Count"`
	Age   int   `bson:"_id"`
}

func (ui *userInfo) QueryCount(username string) ([]QueryCount, error) {
	var (
		list = make([]QueryCount, 0)
		item = QueryCount{}
	)
	pip := bson.A{
		bson.M{"$match": bson.M{"UserName": username}},
		bson.M{"$group": bson.M{"_id": "$Age", "Count": bson.M{"$sum": 1}}},
	}
	cursor, err := ui.GetColl().Aggregate(context.Background(), pip)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		if err := cursor.Decode(&item); err != nil {
			logrus.WithFields(logrus.Fields{"Scenes": "decode error"}).Error(err)
			continue
		}
		list = append(list, item)
	}
	return list, nil
}
