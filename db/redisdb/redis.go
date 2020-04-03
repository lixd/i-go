package redisdb

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"i-go/utils"
	"time"
)

var RedisClient *redis.Client

type redisConf struct {
	Addr               string        `json:"addr"`
	Password           string        `json:"password"`
	DB                 int           `json:"db"`
	MaxRetries         int           `json:"MaxRetries"` // 命令运行失败最大重试次数
	PoolSize           int           `json:"PoolSize"`
	MinIdleConns       int           `json:"MinIdleConns"`
	MaxConnAge         time.Duration `json:"MaxConnAge"`
	PoolTimeout        time.Duration `json:"PoolTimeout"`
	IdleTimeout        time.Duration `json:"IdleTimeout"`
	IdleCheckFrequency time.Duration `json:"IdleCheckFrequency"`
}

func init() {
	defer utils.InitLog("redis")()

	var c redisConf
	if err := viper.UnmarshalKey("redis", &c); err != nil {
		panic(err)
	}
	RedisClient = redis.NewClient(&redis.Options{
		Addr:               c.Addr,
		Password:           c.Password,
		DB:                 c.DB,
		MaxRetries:         c.MaxRetries,
		PoolSize:           c.PoolSize,
		MinIdleConns:       c.MinIdleConns,
		MaxConnAge:         c.MaxConnAge * time.Second,
		PoolTimeout:        c.PoolTimeout * time.Second,
		IdleTimeout:        c.IdleTimeout * time.Second,
		IdleCheckFrequency: c.IdleCheckFrequency * time.Minute,
	})

	fmt.Println(RedisClient)
}

func Release() {
	if RedisClient != nil {
		_ = RedisClient.Close()
		logrus.Info("redis is closed")
	}
}
