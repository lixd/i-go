package stan_pub

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"i-go/mq/nats/constant"
	"i-go/mq/nats/stan/conn"
	"i-go/mq/nats/stan/msghandler"
)

const (
	ErrTimeLimit = 10 // 失败重连临界值
)

var (
	errTime int32 // 失败次数
	config  conn.NATSConfig

	nc  *nats.Conn
	sc  stan.Conn
	err error
)

func init() {
	config := conn.NewQueueNATSConfig(constant.DefaultSubject, constant.DefaultQueue, msghandler.Simple)
	if nc, sc, err = conn.NewConn(config); err != nil {
		panic(err)
	}
	go reConn()
}

// PublishMsg
func PublishMsg(subject string, msg []byte) {
	var (
		GUID  string // 消息发送时返回的序列号
		gLock sync.Mutex
	)
	// ack 回调
	ackHandler := func(guid string, err error) {
		gLock.Lock()
		logrus.Printf("Received ACK for guid %s\n", guid)
		defer gLock.Unlock()
		if err != nil {
			logrus.Fatalf("Error in server ack for guid %s: %v\n", guid, err)
		}
		// 如果两个guid对不上则有问题
		if guid != GUID {
			logrus.Fatalf("Expected a matching guid in ack callback, got %s vs %s\n", guid, guid)
		}
	}

	gLock.Lock()
	GUID, err = sc.PublishAsync(subject, msg, ackHandler)
	if err != nil {
		logrus.Error(err)
		// 失败则增加errTime
		atomic.AddInt32(&errTime, 1)
	}
	gLock.Unlock()
}

// reConn 重连
func reConn() {
	for range time.Tick(time.Second * 30) {
		if errTime < ErrTimeLimit {
			continue
		}
		if nc, sc, err = conn.NewConn(config); err != nil {
			logrus.WithField("scene", "nats reconn").Error(err)
			continue
		}
		// 重连成功归零失败次数
		atomic.AddInt32(&errTime, -errTime)
	}
}

func Release() {
	if sc != nil {
		_ = sc.Close()
	}
	if nc != nil {
		nc.Close()
	}
}
