package pgsqldb

import (
	"fmt"
	"i-go/utils"

	"github.com/go-pg/pg"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var PostgresDB *pg.DB

type PgConf struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	PoolSize int    `json:"poolSize"`
}

func Init() {
	defer utils.InitLog("PostgresSQL")()

	// 0.读取配置文件
	c := parseConf()
	// 1.建立连接
	PostgresDB = newConn(c)
}

func newConn(c *PgConf) *pg.DB {
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	return pg.Connect(&pg.Options{
		User:     c.Username,
		Addr:     addr,
		Password: c.Password,
		Database: c.Database,
		PoolSize: c.PoolSize,
	})
}

func parseConf() *PgConf {
	var c PgConf
	// 0.读取配置文件
	if err := viper.UnmarshalKey("pgsql", &c); err != nil {
		panic(err)
	}
	if c.Host == "" {
		panic("pgsql conf nil")
	}
	return &c
}

func Release() {
	if PostgresDB != nil {
		err := PostgresDB.Close()
		if err != nil {
			logrus.Info("pg close error:", err)
		}
	}
	logrus.Info("pg is closed")
}
