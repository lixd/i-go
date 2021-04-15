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
	// 声明队列 queue
	// params name(队列名) durable(是否持久化) autoDelete(是否自动删除) exclusive(是否排他队列) noWait(指定定义过程是同步还是异步)
	// autoDelete: 在最后一个连接断开后queue会自动删除
	// exclusive: 设置了排外为true的队列只可以在本次的连接中被访问，也就是说在当前连接创建多少个channel访问都没有关系，但是如果是一个新的连接来访问，对不起，不可以
	// x-expires 设置队列有效期 过期后将删除队列(但不保证能及时删除)
	// x-message-ttl 给进入队列的消息设置有效期
	var args = make(amqp.Table)
	args["x-expires"] = 1800000
	args["x-message-ttl"] = 6000
	// x-dead-letter-exchange 指定queue的死信队列为`ExChangeDLX`
	args["x-dead-letter-exchange"] = ExChangeDLX
	// x-dead-letter-routing-key 指定死信消息的新路由键 未指定则使用消息原来的路由键
	args["x-dead-letter-routing-key"] = RoutingKeyDLX
	q, err := ch.QueueDeclare(Queue, false, false, false, false, args)
	qdlx, err := ch.QueueDeclare(QueueDLX, false, false, false, false, args)
	// 交换器 exChange
	// params name(交换器名) type(交换器类型(常用的4种)) durable(是否持久化) autoDelete(是否自动删除(同上 没有剩余bangding时自动删除))
	// internal(是否内置交换器) noWait同上
	// 内置交换器:无法直接发消息到内置交换器，只能通过交换器路由到内置交换器
	// 是否自动删除
	err = ch.ExchangeDeclare(ExChange, amqp.ExchangeFanout, false, false, false, false, nil)
	err = ch.ExchangeDeclare(ExChangeDLX, amqp.ExchangeFanout, false, false, false, false, nil)
	// 将队列绑定到交换器上 bindingKey
	// params name(队列名称) key(bindingKey) exchange(交换器名称) noWait(同上)
	err = ch.QueueBind(q.Name, RoutingKey, ExChange, false, nil)
	err = ch.QueueBind(qdlx.Name, RoutingKey, ExChange, false, nil)

	/*	// 还可以将交换器和交换器绑定 将目标交换器绑定到源交换器上同上指定BindingKey
		// params destination(目标交换器), key(BindingKey), source(源交换器) noWait同上
		ch.ExchangeBind("exDest", "bindingKey", "exSrc",false,nil)*/

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
	/*	// 拉模式消费者
		// 注意:不能再循环里调用Get来代替Consume 会严重影响性能
		msg, ok, err := ch.Get(Queue, false)
		if ok {
			logrus.Printf("Received a message: %s", msg.Body)
		}
	*/
	// 生产消息 routingKey
	for i := 0; i < 11; i++ {

		// params exchange(交换器名称) key(RoutingKey) mandatory(不可达时是否回退给发送者) immediate(无消费者时是否路由到队列) msg(具体的消息)
		// mandatory:为true时 如果交换器无法根据自身类型和路由键找到一个符合条件的队列 那么会将消息返回给生产者 为false则会丢弃消息
		// immediate:为true时 交换器在路由时发现该队列上没有任何消费者，那么将不会把消息存入对列,如果符合条件的队列都没有消费者那么会把消息返回给生产者。
		// msg:参数有点多 DeliveryMode指定消息是否持久化(amqp.Persistent/Transient) Body用于存放真正要发送的消息 Timestamp指定发送时间
		// 		ContentType 消息类型 UserId AppId大概用于指定消息时哪里发出来的吧 还有其它参数一般默认就行。。
		body := "Hello World33!"
		err = ch.Publish(ExChange, RoutingKey, false, false, amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			UserId:       "hello",
			AppId:        "go",
			Timestamp:    time.Now(),
			Body:         []byte(body),
			Expiration:   "1000", // 消息有效期？
		})
	}

	select {}
}
