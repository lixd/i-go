package main

import (
	"log"

	"github.com/Shopify/sarama"
	"i-go/mq/kafka"
	"i-go/mq/kafka/demo/admin"
)

func main() {
	config := sarama.NewConfig()
	ca, err := sarama.NewClusterAdmin([]string{kafka.HOST}, config)
	if err != nil {
		log.Fatal("NewClusterAdmin err: ", err)
	}
	defer ca.Close()
	admin.TopicHelper.List(ca)
	// admin.TopicHelper.Create(ca, kafka.Topic, 1, 1)
	// admin.TopicHelper.Create(ca, kafka.Topic2, 2, 1)
	// admin.TopicHelper.Create(ca, kafka.Topic3, 20, 1)
	// admin.TopicHelper.Delete(ca, kafka.Topic2)
	// admin.TopicHelper.Delete(ca,"new_topic")
	// admin.TopicHelper.Describe(ca, []string{"benchmark"})
	// admin.GroupHelper.List(ca)
	// admin.GroupHelper.Offsets(ca, "g1", "test", []int32{1})
	// err = admin.GroupHelper.Delete(ca, "g1")
	// if err != nil {
	// 	log.Fatal("group Delete err: ", err)
	// }
}
