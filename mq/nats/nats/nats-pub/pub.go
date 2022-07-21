package nats_pub

import (
	"sync/atomic"
	"time"

	"i-go/mq/nats/nats/conn"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

const (
	ErrTimeLimit = 10 // 失败重连临界值
)

var (
	nc      *nats.Conn
	err     error
	errTime int32 // 失败次数
)

func init() {
	nc, err = conn.NewConn()
	if err != nil {
		panic(err)
	}
	go reConn()
}

func PublishMsg(subject string, msg []byte) {
	err := nc.Publish(subject, msg)
	if err != nil {
		logrus.WithField("scene", "nats publish").Error(err)
	}
	_ = nc.Flush()
}

// reConn 重连
func reConn() {
	for range time.Tick(time.Second * 30) {
		if errTime < ErrTimeLimit {
			continue
		}
		if nc, err = conn.NewConn(); err != nil {
			logrus.WithField("scene", "nats reconn").Error(err)
			continue
		}
		// 重连成功归零失败次数
		atomic.AddInt32(&errTime, -errTime)
	}
}
func Release() {
	if nc != nil {
		nc.Close()
	}
}
