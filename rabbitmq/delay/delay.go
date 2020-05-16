package main

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"i-go/core/mq/rabbitmq"
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

var ch = rabbitmq.Channel

// 延迟队列
/*
在定义queue的时候可以通过`x-message-ttl`参数指定进入该队列的消息设置有TTL。

同时在对这个queue添加一个死信交换器和死信队列。这样ttl到了消息就会进入对应的死信队列。

最后消费者订阅死信队列即可达到延迟队列的效果。
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
	// 推模式
	// params name(队列名) consumer(消费者名称或标记) autoAck(是否自动ack) exclusive(是否排他队列) noLocal(暂不支持该参数 只是为了完整性加的) noWait(同时)
	msgs, err := ch.Consume(QueueDLX, strconv.FormatInt(time.Now().UnixNano(), 10), false, false, false, false, nil)
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
			DeliveryMode: amqp.Transient,
			ContentType:  "text/plain",
			UserId:       "guest",
			AppId:        "go",
			Timestamp:    time.Now(),
			Body:         []byte(body),
		})
	}

	select {}
}
