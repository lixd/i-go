package mongodb

import (
	"context"
	"errors"
	"fmt"
	"i-go/utils"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func Init() {
	defer utils.InitLog("MongoDB")()

	c, err := parseConf()
	if err != nil {
		panic(err)
	}
	// 一次初始化多个连接
	newClient(c)
}

func newClient(c *Conf) {
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
			panic(err)
		}
		// 2.初始化客户端
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = MongoClient.Connect(ctx)
		if err != nil {
			logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "conn mongodb"}).Error(err)
			panic(err)
		}
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

func parseConf() (*Conf, error) {
	var c Conf
	if err := viper.UnmarshalKey("mongodb", &c); err != nil {
		return &Conf{}, err
	}
	if c.AppUrl == "" {
		return &Conf{}, errors.New("mongodb conf nil")
	}
	return &c, nil
}

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
