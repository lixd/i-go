package rabbitmq

import (
	"errors"
	"fmt"
	"i-go/utils"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
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

func Init() {
	defer utils.InitLog("RabbitMQ")()
	c, err := parseConf()
	if err != nil {
		panic(err)
	}
	Conn, Channel, err = newConn(c)
	if err != nil {
		panic(err)
	}
}

func parseConf() (*rabbitMQConf, error) {
	var c rabbitMQConf
	if err := viper.UnmarshalKey("rabbitmq", &c); err != nil {
		return nil, err
	}
	if c.Host == "" {
		return nil, errors.New("rabbitmq conf nil")
	}
	return &c, nil
}

func newConn(c *rabbitMQConf) (*amqp.Connection, *amqp.Channel, error) {
	var err error

	// 1.建立连接
	// DSN (Data Source Name)格式: schema://username:password@host:port
	// eg: amqp://guest:guest@192.168.1.111:5672
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%v", c.Username, c.Password, c.Host, c.Port)
	logrus.Info("rabbitmq dsn:", dsn)
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return &amqp.Connection{}, &amqp.Channel{}, err
	}
	// 2.打开通道 信道(基于TCP连接的虚拟连接)
	channel, err := Conn.Channel()
	if err != nil {
		return &amqp.Connection{}, &amqp.Channel{}, err
	}
	return conn, channel, nil
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
