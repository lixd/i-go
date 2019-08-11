package demo

import (
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"log"
)

func SubClient() {
	// 1.连接到默认服务器
	sc, err := GetConn("sub")
	if err != nil {
		logrus.Error(err)
	}
	defer sc.Close()
	subscription, err := sc.Subscribe("illusory", func(msg *stan.Msg) {
		log.Println(msg)
	})
	if err != nil {
		logrus.Error(err)
	}
	defer subscription.Close()
	select {}
}
