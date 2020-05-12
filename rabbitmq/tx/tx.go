package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"i-go/core/mq/rabbitmq"
	"strconv"
	"time"
)

const (
	Queue      = "q-hello"
	ExChange   = "ex-hello"
	BindingKey = "mq.rabbit.hello"
	RoutingKey = "mq.rabbit.hello"
)

// RabbitMQ事务
var (
	conn = rabbitmq.Conn
	ch   = rabbitmq.Channel
)

func main() {
	q, err := ch.QueueDeclare(Queue, false, true, false, false, nil)
	if err != nil {
		logrus.Error("队列定义:", err)
	}
	fmt.Printf("%#v\n", q)
	err = ch.ExchangeDeclare(ExChange, amqp.ExchangeTopic, false, true, false, false, nil)
	if err != nil {
		logrus.Error("交换器定义", err)
	}
	err = ch.QueueBind(q.Name, BindingKey, ExChange, false, nil)
	if err != nil {
		logrus.Error("队列绑定:", err)
	}

	// 消费消息
	go consumer(ch)
	// 生产消息 routingKey
	go publish(ch)

	select {}
}

func consumer(ch *amqp.Channel) {
	// 推模式
	msgs, err := ch.Consume(Queue, strconv.Itoa(int(time.Now().UnixNano())), false, false, false, false, nil)
	if err != nil {
		logrus.Error("消费消息失败", err)
	}
	go func() {
		for msg := range msgs {
			logrus.Printf("Received a message: %s", msg.Body)
			msg.Ack(false)
		}
	}()
}

func publish(ch *amqp.Channel) {
	// 开启事务
	err := ch.Tx()
	if err != nil {
		logrus.Error("事务开启失败:", err)
	}
	body := "Hello World33!"
	// 发送消息
	err = ch.Publish(ExChange, RoutingKey, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		UserId:       "guest",
		AppId:        "go",
		Timestamp:    time.Now(),
		Body:         []byte(body),
	})
	// 根据返回的error判定是回滚还是提交
	if err != nil {
		fmt.Println("send error: ", err)
		err := ch.TxRollback()
		if err != nil {
			logrus.Error("事务回滚失败:", err)
		}
	} else {
		err := ch.TxCommit()
		if err != nil {
			logrus.Error("事务提交失败:", err)
		}
	}
}
