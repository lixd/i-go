package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"i-go/database/mongodb"
	"i-go/gin/demo/model"
)

type DAO interface {
	FindUserByPhone(phone string) (user *model.User, err error)
}
type AccountDAO struct {
	coll *mongo.Collection
}

func (dao *AccountDAO) GetColl() *mongo.Collection {
	if dao.coll == nil {
		dao.coll = mongodb.GetCollection(new(model.User))
	}
	return dao.coll
}
func (dao *AccountDAO) FindUserByPhone(phone string) (user *model.User, err error) {
	filter := bson.M{"phone": phone}
	option := options.FindOne().SetProjection(bson.M{
		"Name":     1,
		"Password": 1})
	one := dao.GetColl().FindOne(context.Background(), filter, option)
	result := model.User{}
	_ = one.Decode(&result)
	return &result, one.Err()
}
