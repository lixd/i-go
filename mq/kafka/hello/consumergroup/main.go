package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
	"i-go/mq/kafka"
)

// 使用 sarama.ConsumerGroup 接口，作为自定义ConsumerGroup
type exampleConsumerGroupHandler struct{}

func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	cg, err := sarama.NewConsumerGroup([]string{kafka.HOST}, kafka.ConsumerGroupID, config)
	if err != nil {
		log.Fatal("NewConsumerGroup err: ", err)
	}
	defer cg.Close()
	handler := exampleConsumerGroupHandler{}
	for {
		err := cg.Consume(context.Background(), []string{kafka.Topic}, handler)
		if err != nil {
			log.Fatal("Consume err: ", err)
		}
	}
}
