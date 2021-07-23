package main

import (
	"time"

	"i-go/mq/kafka"
	"i-go/mq/kafka/demo/producer/sync"
)

func main() {
	topic := kafka.Topic
	// 先启动消费者,保证能消费到后续发送的消息
	// go standalone.SinglePartition(topic)
	sync.Producer(topic, 1000)
	// sleep 等待消费结束
	time.Sleep(time.Second * 10)
}
