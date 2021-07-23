package admin

import (
	"fmt"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
)

/*
	本例展示 使用 ClusterAdmin API 管理 kafka topic
*/

type topic struct{}

var TopicHelper = &topic{}

func (t topic) List(ca sarama.ClusterAdmin) {
	topics, err := ca.ListTopics()
	if err != nil {
		log.Fatal("NewClusterAdmin err: ", err)
	}
	fmt.Println("--------------Topics---------------")
	for k, v := range topics {
		fmt.Printf("name:%s partition:%v replica:%v \n", k, v.NumPartitions, v.ReplicationFactor)
	}
	fmt.Printf("--------------Topics---------------\n\n")
}

func (t topic) Create(ca sarama.ClusterAdmin, topic string, partition, replica int64) error {
	topicDetail := &sarama.TopicDetail{
		NumPartitions:     int32(partition), // 分区数
		ReplicationFactor: int16(replica),   // 备份数
	}
	err := ca.CreateTopic(topic, topicDetail, false) // validateOnly 设置为false才行
	return err
}

func (t topic) Delete(ca sarama.ClusterAdmin, topic string) error {
	return ca.DeleteTopic(topic)
}

// Describe topic 详细信息
func (t topic) Describe(ca sarama.ClusterAdmin, topics []string) {
	describeTopics, err := ca.DescribeTopics(topics)
	if err != nil {
		log.Fatal("DescribeTopics err: ", err)
	}
	for _, topicInfo := range describeTopics {
		fmt.Printf("name:%s IsInternal:%v \n", topicInfo.Name, topicInfo.IsInternal)
		for _, partition := range topicInfo.Partitions {
			fmt.Printf("partition:%+v\n", partition)
		}
	}
}
