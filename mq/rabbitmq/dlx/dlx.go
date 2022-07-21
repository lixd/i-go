package main

import (
	"i-go/core/mq/rabbitmq"
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

var ch = rabbitmq.Channel

//死信队列
/*
定义两个交换器，一个用于业务一个做死信交换器(**两个都是普通交换器**)。

定义一个正常业务队列,定义的时候需要指定一个交换器作为该队列的死信交换器，同时还还可以指定死信消息重新发到死信交换器后需不需要改变routingKey之类的。

然后在定义一个普通队列做为死信队列，绑定到死信交换器上。
*/
func main() {
	// 定义交换器 一个用作业务交换器 一个用作死信交换器
	err := ch.ExchangeDeclare(ExChange, amqp.ExchangeTopic, false, true, false, false, nil)
	err = ch.ExchangeDeclare(ExChangeDLX, amqp.ExchangeFanout, false, true, false, false, nil)

	// 死信队列
	qDxl, err := ch.QueueDeclare(QueueDLX, false, true, false, false, nil)
	if err != nil {
		panic(err)
	}
	// 声明队列 一个用作业务队列 一个用作死信队列

	// 定义正常队列的时候指定一个死信交换器
	var args = make(amqp.Table)
	args["x-message-ttl"] = 1000 // 延迟1秒
	// x-dead-letter-exchange 指定queue的死信队列为`ExChangeDLX`
	args["x-dead-letter-exchange"] = ExChangeDLX
	// x-dead-letter-routing-key 指定死信消息的新路由键 未指定则使用消息原来的路由键
	args["x-dead-letter-routing-key"] = RoutingKeyDLX
	// 声明队列 一个用作业务一个用作死信
	qNormal, err := ch.QueueDeclare(Queue, false, true, false, false, args)
	if err != nil {
		panic(err)
	}

	// 将队列绑定到交换器上
	err = ch.QueueBind(qNormal.Name, RoutingKey, ExChange, false, nil)
	err = ch.QueueBind(qDxl.Name, RoutingKeyDLX, ExChangeDLX, false, nil)
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
			UserId:       "guest",
			AppId:        "go",
			Timestamp:    time.Now(),
			Body:         []byte(body),
			Expiration:   "1000", // 消息有效期？
		})
	}

	select {}
}
