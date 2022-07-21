package mysqldb

import (
	"errors"
	"fmt"

	"i-go/utils"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	_ = MySQL.Callback().Create().After("gorm:after_create").Register(callBackLogName, afterLog)
	_ = MySQL.Callback().Query().After("gorm:after_query").Register(callBackLogName, afterLog)
	_ = MySQL.Callback().Delete().After("gorm:after_delete").Register(callBackLogName, afterLog)
	_ = MySQL.Callback().Update().After("gorm:after_update").Register(callBackLogName, afterLog)
	_ = MySQL.Callback().Row().After("gorm:row").Register(callBackLogName, afterLog)
	_ = MySQL.Callback().Raw().After("gorm:raw").Register(callBackLogName, afterLog)
}

const callBackLogName = "zlog"

func afterLog(db *gorm.DB) {
	err := db.Error
	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	if err != nil {
		return
	}
	fmt.Println("sql:", sql)
}

func parseConf() (*mysqlConf, error) {
	var c mysqlConf
	if err := viper.UnmarshalKey("mysql", &c); err != nil {
		return nil, err
	}
	if c.Host == "" {
		return nil, errors.New("mysql conf nil")
	}
	return &c, nil
}

func newConn(c *mysqlConf) (*gorm.DB, error) {
	// 1.建立连接
	// DSN (Data Source Name)格式: [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	// eg: root:123456@tcp(192.168.100.111:3306)/sampdb?charset=utf8&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		c.Username, c.Password, c.Host, c.Port, c.Database)
	logrus.Info("mysql dsn:", dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
