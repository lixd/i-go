package main

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"strconv"
	"time"
)

const (
	Queue         = "q-hello"
	QueueDLX      = "q-hello-dlx"
	ExChange      = "ex-hello"
	ExChangeDLX   = "ex-hello-dlx"
	RoutingKey    = "mq.rabbit.hello"
	RoutingKeyDLX = "mq.rabbit.hello.dlx"
)

//优先级队列 x-max-priority参数指定优先级
func main() {

	// 建立连接
	conn, err := amqp.Dial("amqp://guest:guest@192.168.100.111:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 打开通道 信道(基于TCP连接的虚拟连接)
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	// 定义交换器 一个用作业务一个用作死信
	err = ch.ExchangeDeclare(ExChange, amqp.ExchangeFanout, false, false, false, false, nil)
	// 声明队列
	var args = make(amqp.Table)
	// x-max-priority 指定队列最大优先级
	args["x-max-priority"] = 5
	qNormal, err := ch.QueueDeclare(QueueDLX, false, false, false, false, args)
	if err != nil {
		panic(err)
	}
	// 将队列绑定到交换器上
	err = ch.QueueBind(qNormal.Name, RoutingKey, ExChange, false, nil)
	// 消费消息
	//推模式
	// params name(队列名) consumer(消费者名称或标记) autoAck(是否自动ack) exclusive(是否排他队列) noLocal(暂不支持该参数 只是为了完整性加的) noWait(同时)
	msgs, err := ch.Consume(Queue, strconv.FormatInt(time.Now().UnixNano(), 10), false, false, false, false, nil)
	go func() {
		for msg := range msgs {
			logrus.Printf("Received a message: %s", msg.Body)
			msg.Ack(false)
		}
	}()

	// 生产消息 routingKey
	for i := 0; i < 11; i++ {
		body := "Hello World33!"
		err = ch.Publish(ExChange, RoutingKey, false, false, amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Timestamp:    time.Now(),
			Body:         []byte(body),
			Priority:     1,
		})
	}

	select {}
}
