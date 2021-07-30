package redisdb

import (
	"time"

	"i-go/utils"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Cli *redis.Client

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

func Init() {
	defer utils.InitLog("Redis")()

	c := parseConf()
	Cli = newClient(c)
}

func Client() *redis.Client {
	if Cli == nil {
		Init()
	}
	return Cli
}

func parseConf() *redisConf {
	var c redisConf
	if err := viper.UnmarshalKey("redis", &c); err != nil {
		panic(err)
	}
	if c.Addr == "" {
		panic("load conf error")
	}
	return &c
}

func newClient(c *redisConf) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:               c.Addr,
		Password:           c.Password,
		DB:                 c.DB,
		MaxRetries:         c.MaxRetries,
		PoolSize:           c.PoolSize,
		MinIdleConns:       c.MinIdleConns,
		MaxConnAge:         c.MaxConnAge,
		PoolTimeout:        c.PoolTimeout,
		IdleTimeout:        c.IdleTimeout,
		IdleCheckFrequency: c.IdleCheckFrequency,
	})
}

func Release() {
	if Cli != nil {
		_ = Cli.Close()
		logrus.Info("redis is closed")
	}
}
