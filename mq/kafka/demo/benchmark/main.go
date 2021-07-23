package main

import (
	"strconv"

	"i-go/mq/kafka"
	"i-go/mq/kafka/demo/consumer/group"
	"i-go/mq/kafka/demo/producer/async"
)

func main() {
	topic := kafka.Topic3
	limit := 100000
	for i := 0; i < 20; i++ {
		go group.ConsumerGroup(topic, "cg2", "CG"+strconv.Itoa(i))
	}
	for i := 0; i < 100; i++ {
		go async.Producer(topic, limit)
	}
	select {}
}
