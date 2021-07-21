package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"i-go/mq/kafka"
)

/*
	本例展示
		0 指定 topic 和 partition， 获得对应的 leader broker
		1 新建 topic
*/

func main() {
	config := sarama.NewConfig()
	client, err := sarama.NewClient([]string{kafka.HOST}, config)
	if err != nil {
		log.Fatal("NewClient err:", err)
	}
	defer client.Close()

	leaderBroker, _ := client.Leader(kafka.Topic, 0)

	topicDetail := &sarama.TopicDetail{}
	topicDetail.NumPartitions = int32(1)                 // 分区数
	topicDetail.ReplicationFactor = int16(1)             // 备份数
	topicDetail.ConfigEntries = make(map[string]*string) // 不知道
	// 创建 topic
	topicDetails := make(map[string]*sarama.TopicDetail)
	topicDetails["new_topic_from_client_test"] = topicDetail
	resp, err := leaderBroker.CreateTopics(&sarama.CreateTopicsRequest{
		TopicDetails: topicDetails,
		Timeout:      time.Second * 15,
	})
	if err != nil {
		panic(err)
	}
	t := resp.TopicErrors
	for key, val := range t {
		fmt.Printf("Key: '%s', Err: %#v, ErrMsg: %#v\n", key, val.Err.Error(), val.ErrMsg)
	}
	client.RefreshMetadata() // 重新获取元数据
	// 删除 topic
	delResp, err := leaderBroker.DeleteTopics(&sarama.DeleteTopicsRequest{
		Timeout: time.Second * 15,
		Topics:  []string{"new_topic_from_client_test"},
	})
	if err != nil {
		panic(err)
	}
	for key, val := range delResp.TopicErrorCodes {
		fmt.Printf("Key: '%s', Err: %#v\n", key, val.Error())
	}
	fmt.Println(client.Topics())
}
