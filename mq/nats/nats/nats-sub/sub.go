package nats_sub

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"i-go/mq/nats/nats/conn"
)

var (
	subs []*nats.Subscription
	nc   *nats.Conn
	err  error
)

func Subscribe(subject string, msgHandler func(msg *nats.Msg)) {
	nc, err = conn.NewConn()
	if err != nil {
		panic(err)
	}
	sub, err := nc.Subscribe(subject, msgHandler)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "subscribe error"}).Error(err)
	}

	err = nc.Flush()

	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "subscribe error"}).Error(err)
	}

	subs = append(subs, sub)
}

func UnSubscribe() {
	for _, sub := range subs {
		_ = sub.Unsubscribe()
	}
	if nc != nil {
		nc.Close()
	}
}
