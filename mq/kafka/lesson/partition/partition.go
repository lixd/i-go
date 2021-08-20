package partition

import (
	"log"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
	"i-go/mq/kafka"
)

type myPartitioner struct {
	partition int32
}

func NewMyPartitioner(topic string) sarama.Partitioner {
	return &myPartitioner{}
}

func (p *myPartitioner) Partition(message *sarama.ProducerMessage, numPartitions int32) (int32, error) {
	if p.partition >= numPartitions {
		p.partition = 0
	}
	ret := p.partition
	p.partition++
	return ret, nil
}

func (p *myPartitioner) RequiresConsistency() bool {
	return false
}

func Producer(topic string, limit int) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = NewMyPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true // 这个默认值就是 true 可以不用手动 赋值

	producer, err := sarama.NewSyncProducer([]string{kafka.HOST}, config)
	if err != nil {
		log.Fatal("NewSyncProducer err:", err)
	}
	defer producer.Close()
	for i := 0; i < limit; i++ {
		str := strconv.Itoa(int(time.Now().UnixNano()))
		msg := &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(str)}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Println("SendMessage err: ", err)
			return
		}
		log.Printf("[Producer] partitionid: %d; offset:%d, value: %s\n", partition, offset, str)
	}
}

/*
生产者
	基本使用 √
	分区机制解析
	生产者压缩算法
消费者
	基本使用
	消费者组
	消费者组重平衡能避免吗
	消费者组消费进度监控
Broker
	位移主题
	无消息丢失配置
	幂等生产者 事务消费者 https://www.jianshu.com/p/f77ade3f41fd
	客户端高级功能 拦截器
原理分析
	Kafka 副本机制
	请求是怎么被处理的
	消费者组重平衡流程解析
	Kafka控制器
	高水位和LeaderEpoch
管理与监控
	AdminClient
	监控工具
	MirrorMaker 跨集群备份
	调优Kafka
*/
