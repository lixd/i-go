package main

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

// 测试备份交换器
const (
	Queue        = "q-hello"
	QueueBack    = "q-hello-back"
	ExChange     = "ex-hello"
	ExChangeBack = "ex-hello-back"
	BindingKey   = "mq.rabbit.hello"
	RoutingKey   = "mq.rabbit.hello"
)

func main() {
	// 1.建立连接
	conn, err := amqp.Dial("amqp://guest:guest@192.168.1.111:5672/")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v \n", conn)
	defer conn.Close()
	// 2.打开通道 信道(基于TCP连接的虚拟连接)
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	// 3.声明队列 queue
	q, err := ch.QueueDeclare(Queue, false, true, false, false, nil)
	q2, err := ch.QueueDeclare(QueueBack, false, true, false, false, nil)
	fmt.Printf("%#v\n", q)
	fmt.Printf("%#v\n", q2)

	// 4.交换器
	// normalExChange
	err = ch.ExchangeDeclare(ExChange, amqp.ExchangeTopic, false, true, false, false, nil)
	// 备份交换器 推荐type为fanout
	args := amqp.Table{"alternate-exchange": ExChangeBack}
	err = ch.ExchangeDeclare(ExChangeBack, amqp.ExchangeFanout, false, true, false, false, args)

	// 5.将队列绑定到交换器上 bindingKey
	err = ch.QueueBind(Queue, BindingKey, ExChange, false, nil)
	err = ch.QueueBind(QueueBack, "", ExChangeBack, false, nil)

	// 6.消费消息
	// 正常队列
	msgs, err := ch.Consume(Queue, "consumerNormal", false, false, false, false, nil)
	go func() {
		for msg := range msgs {
			logrus.Printf("Received a message: %s", msg.Body)
			// multiple为true则会把同一channel上之前接收的消息都Aack或者Nack掉。
			// requeue 表示是否需要把这条消息重新插入队列 再次投递
			msg.Ack(false)
			// msg.Nack(false, true)
			// msg.Reject(true)
		}
	}()
	// 备份队列
	msgs2, err := ch.Consume(QueueBack, "consumerBacK", false, false, false, false, nil)
	go func() {
		for msg := range msgs2 {
			logrus.Printf("back Received a message: %s", msg.Body)
			// multiple为true则会把同一channel上之前接收的消息都Aack或者Nack掉。
			// requeue 表示是否需要把这条消息重新插入队列 再次投递
			msg.Ack(false)
			// msg.Nack(false, true)
			// msg.Reject(true)
		}
	}()
	// 7. 生产消息 routingKey
	for i := 0; i < 11; i++ {
		body := "Hello World33!"
		// 随意指定key
		err = ch.Publish(ExChange, RoutingKey, false, false, amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			UserId:       "guest",
			AppId:        "go",
			Timestamp:    time.Now(),
			Body:         []byte(body),
		})
		if err != nil {
			fmt.Println("send error: ", err)
		} else {
			fmt.Println("send success ", i)
		}
	}
	for i := 0; i < 11; i++ {
		body := "Hello World33!"
		// 随意指定key
		err = ch.Publish(ExChange, "", false, false, amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			UserId:       "guest",
			AppId:        "go",
			Timestamp:    time.Now(),
			Body:         []byte(body),
		})
		if err != nil {
			fmt.Println("send error: ", err)
		} else {
			fmt.Println("send success ", i)
		}
	}
	select {}
}
