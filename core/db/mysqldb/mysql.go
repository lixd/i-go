package mysqldb

import (
	"errors"
	"fmt"
	"i-go/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	MySQL *gorm.DB
)

type mysqlConf struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	Host            string `json:"host"`
	Port            int    `json:"port"`
	Database        string `json:"database"`
	Timeout         string `json:"timeout"`
	TableNamePrefix string `json:"tableNamePrefix"`
}

func Init() {
	defer utils.InitLog("MySQL")()

	c, err := parseConf()
	if err != nil {
		panic(err)
	}

	MySQL, err = newConn(c)
	if err != nil {
		panic(err)
	}

	// 启用Logger，显示详细日志
	MySQL.LogMode(true)

	/*	// 修改默认表名 统一增加前缀
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return c.TableNamePrefix + defaultTableName
		}*/

}

func parseConf() (*mysqlConf, error) {
	var c mysqlConf
	if err := viper.UnmarshalKey("mysql", &c); err != nil {
		return &mysqlConf{}, err
	}
	if c.Host == "" {
		return &mysqlConf{}, errors.New("mysql conf nil")
	}
	return &c, nil
}

func newConn(c *mysqlConf) (*gorm.DB, error) {
	// 1.建立连接
	// DSN (Data Source Name)格式: [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	// eg: root:123456@tcp(192.168.100.111:3306)/sampdb?charset=utf8&parseTime=True&loc=Local&timeout=10s
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		c.Username, c.Password, c.Host, c.Port, c.Database, c.Timeout)
	logrus.Info("mysql dsn:", dsn)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return &gorm.DB{}, err
	}

	return db, nil
}

func Release() {
	if MySQL != nil {
		err := MySQL.Close()
		if err != nil {
			logrus.Info("mysql close error:", err)
		}
	}
	logrus.Info("mysql is closed")
}
