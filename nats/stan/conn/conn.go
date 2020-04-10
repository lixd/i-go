package conn

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"i-go/nats/constant"
	"time"
)

const (
	ErrTimeLimit = 10
)

var (
	errTime int32 // 失败次数
	config  NATSConfig

	err error
	Nc  *nats.Conn
	Sc  stan.Conn
)

type NATSConfig struct {
	ClusterId   string           // 集群id
	ClientId    string           // 客户端id
	QueueName   string           // 队列名
	StanUrl     string           // 监听url
	Subject     string           // 主题
	DurableName string           // 持久name
	MaxInFlight int              // 无确定ack消息时集群最大发送消息量
	AckWait     time.Duration    // ack等待时间
	MsgHandler  func(msg []byte) // 消息处理func
}

func NewQueueNATSConfig(subject, queue string, msgHandler func(msg []byte)) NATSConfig {
	// 随机uuid作为clientId
	uid := subject + "-" + uuid.NewV4().String()
	return NATSConfig{
		ClusterId:   constant.ClusterID,
		ClientId:    uid,
		QueueName:   queue,
		StanUrl:     StanURL(), // 配置文件读取url
		Subject:     subject,
		DurableName: constant.DurableId,
		MaxInFlight: constant.MaxInflight,
		AckWait:     constant.AckWait,
		MsgHandler:  msgHandler,
	}
}
func NewConn(nsc NATSConfig) (nc *nats.Conn, sc stan.Conn, err error) {
	// nats基础连接
	nc, err = nats.Connect(constant.DefaultNatsURL, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "nats conn error"}).Error(err)
		return nil, nil, err
	}
	sc, err = stan.Connect(
		nsc.ClusterId,
		nsc.ClientId,
		stan.NatsConn(nc), // 设置基础nats连接后 关闭streaming后基础的nats连接也可以继续使用 如果没有同时使用streaming和nats则不用设置
		stan.NatsURL(nsc.StanUrl),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			logrus.Fatalf("Connection lost, reason: %v", reason)
		}), // 连接断开时的回调
	)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "nats streaming conn error"}).Error(err)
		return nil, nil, err
	}
	return nc, sc, nil
}

func StanURL() string {
	url := viper.GetString("stanurl")
	if url == "" {
		url = constant.DefaultNatsURL
	}
	return url
}
