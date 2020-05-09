package mysqldb

import (
	_ "i-go/core/conf"

	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"i-go/utils"
)

var MySQL *gorm.DB

type mysqlConf struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Timeout  string `json:"timeout"`
}

func init() {
	defer utils.InitLog("MySQL")()

	conf := readConf()
	newConn(conf)
}

func readConf() *mysqlConf {
	var c mysqlConf
	if err := viper.UnmarshalKey("mysql", &c); err != nil {
		panic(err)
	}
	if c.Host == "" {
		panic("mysql conf nil")
	}
	return &c
}
func newConn(c *mysqlConf) {
	var err error
	// 1.建立连接
	//DSN (Data Source Name)格式: [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	//eg: root:123456@tcp(192.168.100.111:3306)/sampdb?charset=utf8&parseTime=True&loc=Local&timeout=10s
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		c.Username, c.Password, c.Host, c.Port, c.Database, c.Timeout)
	fmt.Println(dsn)
	MySQL, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
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
