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

//持久化
/*
交换器、队列、消息都需要持久化
交换器、队列定义是durable参数为true即可持久化
消息定义时DeliveryMode: amqp.Persistent即为持久化
*/
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
	// 定义交换器
	err = ch.ExchangeDeclare(ExChange, amqp.ExchangeFanout, true, false, false, false, nil)
	// 声明队列
	q, err := ch.QueueDeclare(Queue, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	// 将队列绑定到交换器上
	err = ch.QueueBind(q.Name, RoutingKey, ExChange, false, nil)

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
			DeliveryMode: amqp.Persistent, // 持久化
			ContentType:  "text/plain",
			UserId:       "guest",
			AppId:        "go",
			Timestamp:    time.Now(),
			Body:         []byte(body),
			Expiration:   "1000", // 消息有效期？
		})
	}

	select {}
}
