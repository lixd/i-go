package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
	"i-go/mq/kafka"
)

func main() {
	// demo0()
	demo1()
}

func demo0() {
	config := sarama.NewConfig()
	client, err := sarama.NewClient([]string{kafka.HOST}, config)
	if err != nil {
		log.Fatal("NewClient err:", err)
	}
	defer client.Close()
	consumer, err := sarama.NewConsumer([]string{kafka.HOST}, config)
	if err != nil {
		log.Fatal("NewConsumer err:", err)
	}
	defer consumer.Close()

	// client.GetOffset(kafka.Topic, 0, sarama.OffsetNewest)
	partitionConsumer, err := consumer.ConsumePartition(kafka.Topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal("ConsumePartition err:", err)
	}

	defer partitionConsumer.Close()
	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d msg:%v\n", msg.Offset, string(msg.Value))
			consumed++
			partitionConsumer.HighWaterMarkOffset()
		case <-signals:
			break ConsumerLoop
		}
	}

	log.Printf("Consumed: %d\n", consumed)
}

func demo1() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	client, err := sarama.NewClient([]string{kafka.HOST}, config)
	defer client.Close()
	if err != nil {
		panic(err)
	}

	// consumer, err := sarama.NewConsumerFromClient(client)
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = time.Millisecond * 10
	consumer, err := sarama.NewConsumer([]string{kafka.HOST}, config)
	if err != nil {
		panic(err)
	}

	defer consumer.Close()
	// get partitionId list
	partitions, err := consumer.Partitions(kafka.Topic)
	if err != nil {
		panic(err)
	}
	for _, partitionId := range partitions {
		// create partitionConsumer for every partitionId
		partitionConsumer, err := consumer.ConsumePartition(kafka.Topic, partitionId, sarama.OffsetOldest)
		if err != nil {
			panic(err)
		}
		for message := range partitionConsumer.Messages() {
			value := string(message.Value)
			log.Printf("Partitionid: %d; offset:%d, value: %s\n", message.Partition, message.Offset, value)
		}
		_ = partitionConsumer.Close()
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	select {
	case <-signals:
	}
}
