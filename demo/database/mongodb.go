package database

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"i-go/demo/conf"
	"time"
)

type Conf struct {
	AppUrl        string `json:"appUrl"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	MaxPoolSize   uint16 `json:"maxPoolSize"`
	DBName        string `json:"dbName"`
	AuthMechanism string `json:"authMechanism"`
}

var (
	mongodb *mongo.Database
)

func InitMongoDB() (*mongo.Database, error) {
	err := conf.Init("demo/conf/config.json")
	if err != nil {
		fmt.Printf("viper err=%v \n", err)
	}
	var c Conf
	// 读取配置文件 config.json
	if err := viper.UnmarshalKey("mongo", &c); err != nil {
		logrus.Panic(err)
	}
	appUrl := fmt.Sprintf("mongodb://%s", c.AppUrl)
	// 获取 client
	client, err := mongo.NewClient(options.Client().ApplyURI(appUrl).SetAuth(options.Credential{
		Username:      c.Username,
		Password:      c.Password,
		AuthMechanism: c.AuthMechanism,
		AuthSource:    c.DBName}))
	if err != nil {
		fmt.Printf("mongo.NewClient error=%v", err)
	}
	// 设置 30s 超时
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// 初始化 client
	errC := client.Connect(ctx)
	if errC != nil {
		fmt.Print(errC)
	}
	// 判断服务是否可用
	errP := client.Ping(ctx, readpref.Primary())
	if errP != nil {
		fmt.Print(errP)
	}
	mongodb = client.Database(c.DBName)
	return mongodb, err
}

// 每个 Model 实现该方法 使用不同的collection
type MongoCollection interface {
	GetCollectionName() string
}

func GetCollection(c MongoCollection) *mongo.Collection {
	database, e := InitMongoDB()
	if e != nil {
		return nil
	}
	return database.Collection(c.GetCollectionName())
}
