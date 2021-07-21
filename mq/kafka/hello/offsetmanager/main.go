package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"i-go/mq/kafka"
)

func main() {
	for i := 0; i < 1; i++ {
		demo()
	}
}

func demo() {
	config := sarama.NewConfig()
	config.Consumer.Offsets.AutoCommit.Enable = true              // 开启自动 commit offset 时间间隔
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second // 自动 commit offset 时间间隔
	client, err := sarama.NewClient([]string{kafka.HOST}, config)
	if err != nil {
		log.Fatal("NewClient err: ", err)
	}
	defer client.Close()
	consumer, _ := sarama.NewConsumerFromClient(client)
	offsetManager, _ := sarama.NewOffsetManagerFromClient("111test", client) // 偏移量管理器
	defer offsetManager.Close()
	partitionOffsetManager, _ := offsetManager.ManagePartition(kafka.Topic, kafka.Partition) // 对应分区的偏移量管理器
	defer partitionOffsetManager.Close()
	defer offsetManager.Commit()                         // 手动 commit 上句代码执行后的 Marked Offset
	nextOffset, _ := partitionOffsetManager.NextOffset() // 取得下一消息的偏移量
	fmt.Println("start nextOffset: ", nextOffset)
	pc, _ := consumer.ConsumePartition(kafka.Topic, kafka.Partition, nextOffset)
	defer pc.Close()

	for message := range pc.Messages() {
		value := string(message.Value)
		log.Printf("Partitionid: %d; offset:%d, value: %s\n", message.Partition, message.Offset, value)
		partitionOffsetManager.MarkOffset(message.Offset+1, "modified metadata") // MarkOffset 更新最后消费的 offset
	}
	/*
		NOTE: 相比普通consumer增加了OffsetManager，调用 MarkOffset 手动记录了当前消费的 offset
		sarama 库的自动提交就相当于 offsetManager.Commit() 操作，还是需要手动调用 MarkOffset。
	*/
}
