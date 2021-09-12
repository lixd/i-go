package main

import (
	"time"

	"i-go/mq/kafka"
	"i-go/mq/kafka/demo/consumer/standalone"
	"i-go/mq/kafka/lesson/partition"
)

func main() {
	topic := kafka.TopicPartition
	// 先启动消费者,保证能消费到后续发送的消息
	go standalone.SinglePartition(topic)
	time.Sleep(time.Second)
	partition.Producer(topic, 100)
	// sleep 等待消费结束
	time.Sleep(time.Second * 10)
}
