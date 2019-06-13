package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"i-go/demo/database"
	"i-go/demo/model"
)

type DAO interface {
	FindUserByPhone(phone string) (user *model.User, err error)
	GetColl() *mongo.Collection
	CreateUser(phone, password string) (userId interface{}, err error)
}

type AccountDAO struct {
	coll *mongo.Collection
}

var ADAO = AccountDAO{}

func (dao *AccountDAO) GetColl() *mongo.Collection {
	if dao.coll == nil {
		dao.coll = database.GetCollection(new(model.User))
	}
	return dao.coll
}
func (dao *AccountDAO) FindUserByPhone(phone string) (user *model.User, err error) {
	filter := bson.M{
		"Phone": phone}
	option := options.FindOne().SetProjection(bson.M{
		"Phone":    1,
		"Password": 1})

	one := dao.GetColl().FindOne(context.Background(), filter, option)
	result := model.User{}
	_ = one.Decode(&result)
	return &result, one.Err()
}
func (dao *AccountDAO) CreateUser(phone, password string) (userId interface{}, err error) {
	doc := &model.User{
		Phone:    phone,
		Password: password}
	insertOne, err := dao.GetColl().InsertOne(context.Background(), doc)
	if err != nil {
		return nil, err
	}
	InsertedID := insertOne.InsertedID
	return InsertedID, nil
}
