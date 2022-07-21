package main

import (
	"fmt"
	"strconv"
	"time"

	"i-go/core/mq/rabbitmq"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

const (
	Queue      = "q-hello"
	ExChange   = "ex-hello"
	BindingKey = "mq.rabbit.hello"
	RoutingKey = "mq.rabbit.hello"
)

// confirm模式
/*
对于消息的确认机制略有些蛋疼。因为在发送的时候不可配置发送的消息id，但在接收确认时，消息id是按照自然数递增的，
也就是说发送者需要按照自然数递增的顺序自己维护发送的消息id。
*/
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
	var (
		deliveryMap        = make(map[uint64]*amqp.Publishing)
		deliveryTag uint64 = 1
		ackTag      uint64 = 1
	)
	go publish(ch, deliveryMap, &deliveryTag, &ackTag)

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

func publish(ch *amqp.Channel, deliveryMap map[uint64]*amqp.Publishing, deliveryTag, ackTag *uint64) {
	// 发送消息
	// 切换到confirm模式
	err := ch.Confirm(false)
	if err != nil {
		logrus.Error("切换到confirm模式失败:", err)
	}
	for i := 0; i < 9999; i++ {
		body := "Hello World33!"
		err = doPublish([]byte(body), deliveryMap, deliveryTag)
		if err != nil {
			logrus.Error("切换到confirm模式失败:", err)
		}
	}
	/*	// 为确认最大10条消息 最大100字节 不开全局
		ch.Qos(10, 100, false)*/
	// 调用NotifyPublish 存储确认结果
	notifyChan := make(chan amqp.Confirmation)
	ch.NotifyPublish(notifyChan)
	for v := range notifyChan {
		fmt.Printf("DeliveryTag:%v Ack%v \n", v.DeliveryTag, v.Ack)
		if v.Ack {
			// ack后就从map里删除
			*ackTag++
			delete(deliveryMap, v.DeliveryTag)
		} else {
			// 	未Ack则可以考虑重发等操作
			err = doPublish(deliveryMap[v.DeliveryTag].Body, deliveryMap, deliveryTag)
			if err != nil {
				logrus.Error("do push:", err)
			}
		}
	}
}

// doPublish 发送消息
func doPublish(body []byte, deliveryMap map[uint64]*amqp.Publishing, deliveryTag *uint64) error {
	msg := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		UserId:       "guest",
		AppId:        "go",
		Timestamp:    time.Now(),
		Body:         body,
	}
	deliveryMap[*deliveryTag] = &msg
	err := ch.Publish(ExChange, RoutingKey, false, false, msg)
	if err != nil {
		fmt.Println("send error: ", err)
	} else {
		// 确认结果中DeliveryTag是递增的 但是发送的时候获取不到这个值 只能自己维护。。
		*deliveryTag++
	}
	return err
}
