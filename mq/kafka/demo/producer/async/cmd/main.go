package main

import (
	"time"

	"i-go/mq/kafka"
	"i-go/mq/kafka/demo/consumer/standalone"
	"i-go/mq/kafka/demo/producer/async"
)

func main() {
	topic := kafka.Topic
	go standalone.SinglePartition(topic)
	time.Sleep(time.Millisecond * 100) // 延迟，让consumer启动后再启动生产者
	async.Producer(topic, 100)

	time.Sleep(time.Second * 10)
}
