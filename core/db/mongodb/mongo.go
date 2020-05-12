package mongodb

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"i-go/utils"
	"time"
)

var TestDB *mongo.Database
var TestClient *mongo.Client

var JobDB *mongo.Database
var JobClient *mongo.Client

type Conf struct {
	AppUrl        string            `json:"appUrl"`
	Username      string            `json:"username"`
	Password      string            `json:"password"`
	MaxPoolSize   uint16            `json:"maxPoolSize"`
	DBS           map[string]string `json:"dbs"`
	AuthMechanism string            `json:"authMechanism"`
}

func init() {
	defer utils.InitLog("mongodb")()

	var c Conf

	// 0.读取配置文件
	if err := viper.UnmarshalKey("mongo", &c); err != nil {
		panic(err)
	}
	// 通过这个一次初始化多个连接
	for key, name := range c.DBS {
		appUrl := fmt.Sprintf("mongodb://%s", c.AppUrl)
		// 1. 获取客户端连接
		MongoClient, err := mongo.NewClient(
			options.Client().ApplyURI(appUrl).
				SetAuth(options.Credential{
					Username: c.Username,
					Password: c.Password,
					// AuthMechanism: c.AuthMechanism0,
					AuthSource: name}).
				SetMaxPoolSize(c.MaxPoolSize),
		)
		if err != nil {
			panic("conn mongodb error")
		}
		// 2.初始化客户端
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = MongoClient.Connect(ctx)
		switch key {
		case "test":
			TestDB = MongoClient.Database(name)
		case "job":
			JobDB = MongoClient.Database(name)
		}
		switch key {
		case "test":
			TestClient = MongoClient
		case "job":
			JobClient = MongoClient
		}
	}
}

type MongoCollection interface {
	GetCollectionName() string
}

func GetTestCollection(c MongoCollection) *mongo.Collection {
	return TestDB.Collection(c.GetCollectionName())
}
func GetJobCollection(c MongoCollection) *mongo.Collection {
	return JobDB.Collection(c.GetCollectionName())
}

const timeout = time.Second * 3

func Release() {
	if TestClient != nil {
		ctx, _ := context.WithTimeout(context.Background(), timeout)
		_ = TestClient.Disconnect(ctx)
	}
	if JobClient != nil {
		ctx, _ := context.WithTimeout(context.Background(), timeout)
		_ = JobClient.Disconnect(ctx)
	}
	logrus.Info("mongodb is closed")
}
