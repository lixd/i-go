package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

const (
	Queue        = "q-hello"
	Queue2       = "q-hello2"
	ExChange     = "ex-hello"
	ExChangeBack = "ex-hello-back"
	BindingKey   = "mq.rabbit.hello"
	RoutingKey   = "mq.rabbit.hello"
)

func main() {

	// 建立连接
	conn, err := amqp.Dial("amqp://guest:guest@47.93.123.142:5672/")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v \n", conn)
	defer conn.Close()
	// 打开通道 信道(基于TCP连接的虚拟连接)
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	// 声明队列 queue
	// params name(队列名) durable(是否持久化) autoDelete(是否自动删除) exclusive(是否排他队列) noWait(指定定义过程是同步还是异步)
	// autoDelete: 在最后一个连接断开后queue会自动删除
	// exclusive: 设置了排外为true的队列只可以在本次的连接中被访问，也就是说在当前连接创建多少个channel访问都没有关系，但是如果是一个新的连接来访问，对不起，不可以
	q, err := ch.QueueDeclare(Queue, false, true, false, false, nil)
	fmt.Printf("%#v\n", q)
	// 交换器 exChange
	// params name(交换器名) type(交换器类型(常用的4种)) durable(是否持久化) autoDelete(是否自动删除(同上 没有剩余bangding时自动删除))
	// internal(是否内置交换器) noWait同上
	// 内置交换器:无法直接发消息到内置交换器，只能通过交换器路由到内置交换器
	// 是否自动删除
	err = ch.ExchangeDeclare(ExChange, amqp.ExchangeFanout, false, true, false, false, nil)
	// 将队列绑定到交换器上 bindingKey
	// params name(队列名称) key(bindingKey) exchange(交换器名称) noWait(同上)
	err = ch.QueueBind(q.Name, BindingKey, ExChange, false, nil)

	/*	// 还可以将交换器和交换器绑定 将目标交换器绑定到源交换器上同上指定BindingKey
		// params destination(目标交换器), key(BindingKey), source(源交换器) noWait同上
		ch.ExchangeBind("exDest", "bindingKey", "exSrc",false,nil)*/
	// 消费消息
	for i := 0; i < 1; i++ {
		time.Sleep(time.Millisecond * 1)
		go consumer(ch)
	}
	/*	// 拉模式消费者
		// 注意:不能再循环里调用Get来代替Consume 会严重影响性能
		msg, ok, err := ch.Get(Queue, false)
		if ok {
			logrus.Printf("Received a message: %s", msg.Body)
		}
	*/

	// 生产消息 routingKey
	for i := 0; i < 1; i++ {
		go publish(ch)
	}

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
			// logrus.Printf("Received a message: %s", msg.Body)
			// multiple为true则会把同一channel上之前接收的消息都Aack或者Nack掉。
			// requeue 表示是否需要把这条消息重新插入队列 再次投递
			msg.Ack(false)
			// msg.Nack(false, true)
			// msg.Reject(true)
		}
	}()
}

func publish(ch *amqp.Channel) {
	for {
		time.Sleep(time.Nanosecond * 1)
		// params exchange(交换器名称) key(RoutingKey) mandatory(不可达时是否回退给发送者) immediate(无消费者时是否路由到队列) msg(具体的消息)
		// mandatory(golang客户端这个参数好像没用？):为true时 如果交换器无法根据自身类型和路由键找到一个符合条件的队列 那么会将消息返回给生产者 为false则会丢弃消息
		// immediate(3.0后已取消该参数):为true时 交换器在路由时发现该队列上没有任何消费者，那么将不会把消息存入对列,如果符合条件的队列都没有消费者那么会把消息返回给生产者。
		// msg:参数有点多 DeliveryMode指定消息是否持久化(amqp.Persistent/Transient) Body用于存放真正要发送的消息 Timestamp指定发送时间
		// 		ContentType 消息类型 UserId就是创建连接时指定的userId 还有其它参数一般默认就行。。
		body := "Hello World33! Hello World33! Hello World33! Hello World33! Hello World33! Hello World33! Hello World33!"
		// operation basic.publish caused a connection exception not_implemented: "immediate=true"
		err := ch.Publish(ExChange, RoutingKey, false, false, amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			UserId:       "guest", // 需要和建立连接时的用户相同
			AppId:        "go",
			Timestamp:    time.Now(),
			Body:         []byte(body),
		})
		if err != nil {
			fmt.Println("send error: ", err)
			continue
		}
	}
}
