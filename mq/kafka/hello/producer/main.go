package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
	"i-go/mq/kafka"
)

func main() {
	producer, err := sarama.NewSyncProducer([]string{kafka.HOST}, nil)
	if err != nil {
		log.Fatal("NewSyncProducer err:", err)
	}
	defer producer.Close()
	// producer.Input() <- &sarama.ProducerMessage{Topic: "my_topic", Key: nil, Value: sarama.StringEncoder("testing 123")}
	for {
		// time.Sleep(time.Second)
		str := strconv.Itoa(int(time.Now().Unix()))
		msg := &sarama.ProducerMessage{Topic: kafka.Topic, Key: nil, Value: sarama.StringEncoder(str)}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Println("SendMessage err: ", err)
			return
		}
		fmt.Printf("msg:%v partition:%v offset:%v\n", str, partition, offset)
	}
}
