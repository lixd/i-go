package db

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"i-go/utils"
)

var PostgresDB *pg.DB

type Conf struct {
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func init() {
	defer utils.InitLog("PostgresSQL")()

	var c Conf

	// 0.读取配置文件
	if err := viper.UnmarshalKey("pgsql", &c); err != nil {
		panic(err)
	}

	addr := fmt.Sprintf("%s:%s", c.Addr, c.Port)
	PostgresDB = pg.Connect(&pg.Options{
		User:     c.Username,
		Addr:     addr,
		Password: c.Password,
		Database: c.Database,
	})
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
