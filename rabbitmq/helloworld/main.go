package main

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"time"
)

func main() {
	// 建立连接
	conn, err := amqp.Dial("amqp://guest:guest@192.168.100.111:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 打开通道
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	// 声明队列
	q, err := ch.QueueDeclare("q-hello", false, false, false, false, nil)
	// 交换器
	err = ch.ExchangeDeclare("ex-demo", amqp.ExchangeFanout, false, false, false, false, nil)
	// 队列绑定到交换器上
	err = ch.QueueBind(q.Name, "", "ex-demo", false, nil)

	// 消费消息
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	go func() {
		for msg := range msgs {
			logrus.Printf("Received a message: %s", msg.Body)
			msg.Ack(false)
		}
	}()
	// 生产消息
	for i := 0; i < 11; i++ {
		body := "Hello World33!"
		err = ch.Publish("", q.Name, false, false, amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Timestamp:    time.Now(),
			Body:         []byte(body),
		})
	}

	select {}
}
