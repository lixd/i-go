package admin

import (
	"fmt"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
)

/*
	本例展示 使用 ClusterAdmin API 管理 kafka
	包括 topic、broker、等管理
*/

type group struct{}

var GroupHelper = &group{}

func (g group) List(ca sarama.ClusterAdmin) {
	topics, err := ca.ListConsumerGroups()
	if err != nil {
		log.Fatal("NewClusterAdmin err: ", err)
	}
	fmt.Println("---------------ConsumerGroups--------------")
	for k, v := range topics {
		fmt.Printf("name:%s protocolType:%s \n", k, v)
	}
	fmt.Printf("--------------ConsumerGroups---------------\n\n")
}

func (g group) Offsets(ca sarama.ClusterAdmin, group, topic string, partition []int32) {
	detail := map[string][]int32{
		topic: partition,
	}
	resp, err := ca.ListConsumerGroupOffsets(group, detail) // validateOnly 设置为false才行
	if err != nil {
		log.Fatal("NewClusterAdmin err: ", err)
	}
	for k, v := range resp.Blocks {
		for kk, vv := range v {
			fmt.Printf("key:%s kk:%+v offset:%v metadata:%s LeaderEpoch:%v err:%v\n", k, kk, vv.Offset, vv.Metadata, vv.LeaderEpoch, vv.Err)
		}
	}
}

func (g group) Delete(ca sarama.ClusterAdmin, group string) error {
	return ca.DeleteConsumerGroup(group)
}
