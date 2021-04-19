package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserInfo struct {
	ID         primitive.ObjectID `bson:"_id"`
	UserName   string             `bson:"UserName"`
	Password   string             `bson:"Password"`
	Age        int                `bson:"Age"`
	Phone      string             `bson:"Phone"`
	Hobby      []string           `bson:"Hobby"`
	CreateTime int64              `bson:"CreateTime"`
	UpdateTime int64              `bson:"UpdateTime"`
}

func (*UserInfo) GetCollectionName() string {
	return "UserInfo"
}

type UserInfoReq struct {
	ID       string     `bson:"_id"`
	UserName string     `bson:"UserName"`
	Password string     `bson:"Password"`
	Age      int        `bson:"Age"`
	Phone    string     `bson:"Phone"`
	Hobby    []string   `bson:"Hobby"`
	Reselect []Reselect `bson:"Reselect"`
}

type Reselect struct {
	Node   string `bson:"Node"`
	Status int    `bson:"Status"`
}
