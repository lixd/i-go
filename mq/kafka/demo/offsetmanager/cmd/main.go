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
	async.Producer(topic, 100)
	time.Sleep(time.Second * 10)
}
