package request_reply

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"i-go/mq/nats/nats/conn"
)

func Subscribe(subject string, buildHandler func(nc *nats.Conn) func(msg *nats.Msg), stopChan <-chan struct{}) {
	nc, err = conn.NewConn()
	if err != nil {
		panic(err)
	}
	sub, err := nc.Subscribe(subject, buildHandler(nc))
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "subscribe error"}).Error(err)
	}
	go func() {
		<-stopChan
		_ = sub.Unsubscribe()
		nc.Close()
	}()
}

func demoHandler(nc *nats.Conn) func(msg *nats.Msg) {
	return func(msg *nats.Msg) {
		fmt.Printf("subject:%s msg:%s\n", msg.Subject, string(msg.Data))
		_ = nc.Publish(msg.Reply, msg.Data)
	}
}
