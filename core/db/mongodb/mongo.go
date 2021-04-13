package mongodb

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"i-go/utils"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// MongoDB Cloud URL
	stdURL = "mongodb+srv://<username>:<password>@clusterfree.p7rd5.mongodb.net/<dbname>?retryWrites=true&w=majority"
)

var (
	TestDB     *mongo.Database
	TestClient *mongo.Client

	JobDB     *mongo.Database
	JobClient *mongo.Client

	XDB     *mongo.Database
	XClient *mongo.Client
)

type Conf struct {
	AppUrl        string            `json:"appUrl"`
	Username      string            `json:"username"`
	Password      string            `json:"password"`
	MaxPoolSize   uint64            `json:"maxPoolSize"`
	MinPoolSize   uint64            `json:"minPoolSize"`
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
	// newClientCluster(c)
}

func newClientCluster(c *Conf) {
	URLUserName := strings.ReplaceAll(stdURL, "<username>", c.Username)
	URLPwd := strings.ReplaceAll(URLUserName, "<password>", c.Password)
	for key, name := range c.DBS {
		URLFull := strings.ReplaceAll(URLPwd, "<dbname>", name)
		fmt.Println("mongodb URL: ", URLFull)
		// appUrl := "mongodb+srv://17x:mongodb12345@clusterfree.p7rd5.mongodb.net/17x?retryWrites=true&w=majority"
		// 1. 获取客户端连接
		MongoClient, err := mongo.NewClient(
			options.Client().
				ApplyURI(URLFull).
				SetMaxPoolSize(c.MaxPoolSize).
				SetMinPoolSize(c.MinPoolSize))
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
			TestClient = MongoClient
		case "job":
			JobDB = MongoClient.Database(name)
			JobClient = MongoClient
		case "17x":
			XDB = MongoClient.Database(name)
			XClient = MongoClient
		}
	}
}
func newClient(c *Conf) {
	for key, name := range c.DBS {
		appUrl := fmt.Sprintf("mongodb://%s", c.AppUrl)
		// 1. 获取客户端连接
		MongoClient, err := mongo.NewClient(
			options.Client().ApplyURI(appUrl).
				// SetAuth(options.Credential{
				// 	Username:      c.Username,
				// 	Password:      c.Password,
				// 	AuthMechanism: c.AuthMechanism,
				// 	AuthSource:    name}).
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
			TestClient = MongoClient
		case "job":
			JobDB = MongoClient.Database(name)
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
