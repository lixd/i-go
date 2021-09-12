package kafka

const (
	HOST             = "123.57.236.125:9092"
	Topic            = "standAlone"
	Topic2           = "consumerGroup"
	Topic3           = "benchmark"
	TopicPartition   = "partition"
	TopicCompression = "compression"
	DefaultPartition = 0
	ConsumerGroupID  = "cgID"
	ConsumerGroupID2 = "cg2"
)

/*
生产者
	基本使用 √
	分区机制解析 √
	生产者压缩算法 √
消费者
	基本使用 √
	消费者组 √
	消费者组重平衡能避免吗 √
	消费者组消费进度监控
Broker
	位移主题
	幂等生产者 事务消费者 https://www.jianshu.com/p/f77ade3f41fd
	客户端高级功能 拦截器
原理分析
	Kafka 副本机制
	请求是怎么被处理的
	消费者组重平衡流程解析
	Kafka控制器
	高水位和LeaderEpoch
其他
	无消息丢失配置 √
	exactlyOnce 语义 √
管理与监控
	AdminClient
	监控工具
	MirrorMaker 跨集群备份
	调优Kafka
*/
