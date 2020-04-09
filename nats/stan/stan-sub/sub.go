package stan_sub

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"i-go/nats/stan/conn"
)

var (
	subs []stan.Subscription
	nc   *nats.Conn
	sc   stan.Conn
	err  error
)

func StartMany(count int, subject, queue string, msgHandler func(msg []byte)) {
	// 开启多个订阅
	//NATS内置负载均衡 虽然有多个Subscribe,但消息来时只会选择其中一个来消费。
	// Although queue groups have multiple subscribers, each message is consumed by only one
	for i := 0; i < count; i++ {
		nsc := conn.NewQueueNATSConfig(subject, queue, msgHandler)
		go queueSubscribe(nsc)
	}
}

// queueSubscribe 队列订阅
func queueSubscribe(nsc conn.NATSConfig) {
	var (
		sub stan.Subscription
	)
	nc, sc, err = conn.NewConn(nsc)
	if err != nil {
		panic(err)
	}

	sub, err = sc.QueueSubscribe(nsc.Subject, nsc.QueueName, func(msg *stan.Msg) {
		if msg.Ack() == nil {
			nsc.MsgHandler(msg.Data)
		}
	}, stan.DurableName(nsc.DurableName),
		stan.MaxInflight(nsc.MaxInFlight),
		stan.SetManualAckMode(), //开启手动确认模式 一旦你配置为手动发送ACK，你必须显式调用 NATS Streaming 消息的函数Ack。 即上边msgHandler中手动调用msg.Ack()
		stan.AckWait(nsc.AckWait))

	if err != nil {
		nc.Close()
		_ = sc.Close()
		logrus.WithFields(logrus.Fields{"Scenes": "nats QueueSubscribe error"}).Error(err)
		return
	}

	subs = append(subs, sub)
}

// Unsubscribe 取消所有订阅
func Unsubscribe() {
	for _, sub := range subs {
		_ = sub.Unsubscribe()
	}
	if sc != nil {
		_ = sc.Close()
	}
	if nc != nil {
		nc.Close()
	}
}
