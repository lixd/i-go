package main

import (
	"time"

	"i-go/mq/kafka"
	"i-go/mq/kafka/demo/consumer/standalone"
	"i-go/mq/kafka/demo/producer/sync"
)

// 测试 独立消费者 先启动消费者再启动生产者
func main() {
	topic := kafka.Topic
	go standalone.SinglePartition(topic)
	// go standalone.Partitions(topic)
	time.Sleep(time.Millisecond * 100)
	sync.Producer(topic, 100)
	time.Sleep(time.Second * 10)
}
