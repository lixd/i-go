package main

import (
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

const (
	Queue         = "q-hello"
	QueueDLX      = "q-hello-dlx"
	ExChange      = "ex-hello"
	ExChangeDLX   = "ex-hello-dlx"
	RoutingKey    = "mq.rabbit.hello"
	RoutingKeyDLX = "mq.rabbit.hello.dlx"
)

// 惰性队列
/*
默认将消息全存储到磁盘上,
定义时增加参数 queue-mode=lazy 即可
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
	err = ch.ExchangeDeclare(ExChange, amqp.ExchangeFanout, false, false, false, false, nil)
	// 声明队列
	// 惰性队列需要增加参数
	var args = make(amqp.Table)
	// queue-mode=lazy 指定queue为惰性队列
	args["queue-mode"] = "lazy"
	qLazy, err := ch.QueueDeclare(QueueDLX, false, false, false, false, args)
	if err != nil {
		panic(err)
	}
	// 将队列绑定到交换器上
	err = ch.QueueBind(qLazy.Name, RoutingKey, ExChange, false, nil)

	// 消费消息
	// 推模式
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
			UserId:       "guest",
			AppId:        "go",
			Timestamp:    time.Now(),
			Body:         []byte(body),
			Expiration:   "1000", // 消息有效期？
		})
	}

	select {}
}
