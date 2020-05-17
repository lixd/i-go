package rabbitmq

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"i-go/core/conf"
	"i-go/utils"
)

var (
	Conn    *amqp.Connection
	Channel *amqp.Channel
)

type rabbitMQConf struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func init() {
	defer utils.InitLog("RabbitMQ")()
	conf.Init("conf/config.json")
	c := readConf()
	newConn(c)
}

func readConf() *rabbitMQConf {
	var c rabbitMQConf
	if err := viper.UnmarshalKey("rabbitmq", &c); err != nil {
		panic(err)
	}
	if c.Host == "" {
		panic("rabbitmq conf nil")
	}
	return &c
}
func newConn(c *rabbitMQConf) {
	var err error

	// 1.建立连接
	// DSN (Data Source Name)格式: schema://username:password@host:port
	// eg: amqp://guest:guest@192.168.1.111:5672
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%v", c.Username, c.Password, c.Host, c.Port)
	logrus.Info("rabbitmq dsn:", dsn)
	Conn, err = amqp.Dial(dsn)
	if err != nil {
		panic(err)
	}
	// 2.打开通道 信道(基于TCP连接的虚拟连接)
	Channel, err = Conn.Channel()
	if err != nil {
		panic(err)
	}
}

func Release() {
	if Channel != nil {
		err := Channel.Close()
		if err != nil {
			logrus.Info("rabbitmq channel close error:", err)
		}
	}
	if Conn != nil {
		err := Conn.Close()
		if err != nil {
			logrus.Info("mysql close error:", err)
		}
	}
	logrus.Info("rabbitmq is closed")
}
