package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

type exampleConsumerGroupHandler struct{}

func (exampleConsumerGroupHandler) Setup(sess sarama.ConsumerGroupSession) error {
	// sess.ResetOffset("test_topic_1",0,10,"")
	// sess.Commit()
	return nil
}
func (exampleConsumerGroupHandler) Cleanup(sess sarama.ConsumerGroupSession) error {

	return nil

}
func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	fmt.Println(claim.InitialOffset())
	// fmt.Println("highwater",claim.HighWaterMarkOffset())

	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		sess.MarkMessage(msg, "")
		fmt.Println(string(msg.Value))
	}
	return nil
}

var kafkaClient sarama.Client
var kafkaClientFirst sync.Once

func GetLwBotKafkaClient() (sarama.Client, error) {
	config := sarama.NewConfig()
	config.Net.ReadTimeout = 180 * time.Second
	config.Consumer.Return.Errors = true
	// config.Consumer.Offsets.AutoCommit.Enable=false
	var err error
	brokers := []string{"119.3.231.213:9092"}
	kafkaClientFirst.Do(func() {
		kafkaClient, err = sarama.NewClient(brokers, config)

	})
	topics, _ := kafkaClient.Topics()
	fmt.Println(topics)
	if err != nil {
		return nil, err
	}
	return kafkaClient, nil
}

func main() {
	fmt.Println("开始运行")
	config := sarama.NewConfig()
	config.Version = sarama.MaxVersion
	group, err := sarama.NewConsumerGroup([]string{"119.3.231.213:9092"}, "test_cg_1", config)

	if err != nil {
		panic(err)
	}
	defer func() { _ = group.Close() }()

	// Track errors
	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()

	ctx := context.Background()
	for {
		topics := []string{"test_topic_1"}
		handler := exampleConsumerGroupHandler{}

		// `Consume` should be called inside an infinite loop, when a
		// server-side rebalance happens, the consumer session will need to be
		// recreated to get the new claims
		err := group.Consume(ctx, topics, handler)
		if err != nil {
			panic(err)
		}
	}
}
