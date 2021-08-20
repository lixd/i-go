package main

import (
	"time"

	"i-go/mq/kafka"
	"i-go/mq/kafka/demo/offsetmanager"
	"i-go/mq/kafka/demo/producer/async"
)

func main() {
	topic := kafka.Topic
	go offsetmanager.OffsetManager(topic)
	time.Sleep(time.Second) // sleep 让 consumer 先启动
	async.Producer(topic, 100)
	time.Sleep(time.Second * 10)
}
