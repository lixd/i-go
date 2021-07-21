package main

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

var kafkaClient sarama.Client
var kafkaClientFirst sync.Once

func GetLwBotKafkaClient() (sarama.Client, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0
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
	c, _ := GetLwBotKafkaClient()

	// admin,err:=sarama.NewClusterAdminFromClient(c)
	// if err != nil {
	//   fmt.Println(err)
	// }
	//
	// err = admin.DeleteConsumerGroup("test_cg_1")
	// if err != nil {
	//    fmt.Println(err)
	// }

	of, err := sarama.NewOffsetManagerFromClient("test_cg_1", c)
	if err != nil {
		fmt.Println(err)
		return
	}

	pom, err := of.ManagePartition("test_topic_1", 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取最早的记录
	expected, err := c.GetOffset("test_topic_1", 0, sarama.OffsetNewest)

	if err != nil {
		fmt.Println(err)
		return
	}

	pom.ResetOffset(expected, "modified_meta")

	actual, meta := pom.NextOffset()
	fmt.Println(actual, meta)
	if actual != expected {
		fmt.Println("Expected offset %v. Actual: %v", expected, actual)
	}
	if meta != "modified_meta" {
		fmt.Println("Expected metadata \"modified_meta\". Actual: %q", meta)
	}
	// pom.AsyncClose()
	pom.Close()
	of.Close()

	// 删除记录
	// req := &sarama.DeleteRecordsRequest{
	//  Topics: map[string]*sarama.DeleteRecordsRequestTopic{
	//      "test_topic_1": {
	//          PartitionOffsets: map[int32]int64{
	//              0: 2,
	//
	//          },
	//      },
	//      "other": {},
	//  },
	//  Timeout: 100 * time.Millisecond,
	// }
	// //
	// b,err:=c.Broker(0)
	// fmt.Println(b.Addr())
	// if err!=nil{
	//   fmt.Println(err)
	// }
	// ////// fmt.Println(b.)
	// r,err:=b.DeleteRecords(req)
	// if err!=nil{
	//   fmt.Println(err)
	// }
	// fmt.Println(r.Topics)

	// fmt.Println(c.Topics())

	// 删除topic
	// req:= &sarama.DeleteTopicsRequest{
	//   Version: 1,
	//   Topics:  []string{"test_topic_2"},
	//   Timeout: 100 * time.Millisecond,
	// }
	// r,err:=b.DeleteTopics(req)
	// if err!=nil{
	//    fmt.Println(err)
	// }
	// fmt.Println(r.TopicErrorCodes)
	// time.Sleep(5*time.Second)
	// fmt.Println(c.Topics())

	// 创建topic
	// retention := "-1"
	// req := &sarama.CreateTopicsRequest{
	//    TopicDetails: map[string]*sarama.TopicDetail{
	//        "test_topic_3": {
	//            NumPartitions:     -1,
	//            ReplicationFactor: -1,
	//            ReplicaAssignment: map[int32][]int32{
	//                0: {0, 1, 2},
	//            },
	//            ConfigEntries: map[string]*string{
	//                "retention.ms": &retention,
	//            },
	//        },
	//    },
	//    Timeout: 100 * time.Millisecond,
	// }
	// //req.Version=1
	// //req.ValidateOnly=true
	//
	// b,err:=c.Broker(0)
	// if err!=nil{
	//   // return nil
	// }
	// rep,err:=b.CreateTopics(req)
	// if err!=nil{
	//   // return err
	// }
	// fmt.Println(rep)

	fmt.Println(c.Topics())
	c.Close()
	fmt.Println(c.Closed())
}
