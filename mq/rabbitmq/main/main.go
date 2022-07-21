package main

import (
	"fmt"

	"i-go/mq/rabbitmq"
)

type TestPro struct {
	msgContent string
}

// 实现发送者
func (t *TestPro) MsgContent() string {
	return t.msgContent
}

// 实现接收者
func (t *TestPro) Consumer(dataByte []byte) error {
	fmt.Println(string(dataByte))
	return nil
}

func main() {
	msg := "这是测试任务"
	t := &TestPro{
		msg,
	}
	queueExchange := &rabbitmq.QueueExchange{
		QuName: "test.rabbit",
		RtKey:  "rabbit.key",
		ExName: "test.rabbit.mq",
		ExType: "direct",
	}
	mq := rabbitmq.New(queueExchange)
	mq.RegisterProducer(t)
	mq.RegisterReceiver(t)
	mq.Start()
}
