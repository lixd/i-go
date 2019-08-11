package demo

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"log"
)

func PublisMsg() {
	sc, err := GetConn("pub")
	if err != nil {
		logrus.Error(err)
	}
	defer sc.Close()

	ackHandler := func(ackedNuid string, err error) {
		log.Println("ackedNuid: ", ackedNuid)
	}

	for i := 0; i < 10; i++ {
		_, err := sc.PublishAsync("illusory", []byte("Hello World"), ackHandler)
		if err != nil {
			logrus.Error(err)
		}
	}
}

func GetConn(clientID string) (stan.Conn, error) {
	// 1.连接到默认服务器
	sc, err := stan.Connect("test-cluster", clientID,
		stan.NatsURL(nats.DefaultURL))
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return sc, nil
}
