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

func (ui *userInfo) Upsert(req *model.UserInfoReq) error {
	var filter bson.M
	if len(req.ID) != 0 {
		objId, err := primitive.ObjectIDFromHex(req.ID)
		if err != nil {
			return err
		}
		filter = bson.M{
			"_id": objId,
		}
	}

	update := bson.M{
		"$setOnInsert": bson.M{
			"CreateTime": time.Now().Unix(),
		},
		"$set": bson.M{
			"UserName":   req.UserName,
			"Password":   req.Password,
			"Age":        req.Age,
			"Phone":      req.Phone,
			"UpdateTime": time.Now().Unix(),
		},
	}
	opts := options.Update().SetUpsert(true)
	_, err := ui.GetColl().UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		logrus.Error(err)
	}
	return nil
}
