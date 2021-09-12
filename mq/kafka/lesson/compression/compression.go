package compression

import (
	"log"
	"strings"

	"github.com/Shopify/sarama"
	"i-go/mq/kafka"
)

var defaultMsg = strings.Repeat("Golang", 1000)

func Producer(topic string, limit int) {
	config := sarama.NewConfig()
	// config.Producer.Compression = sarama.CompressionGZIP
	// config.Producer.CompressionLevel = gzip.BestCompression
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true // 这个默认值就是 true 可以不用手动 赋值

	producer, err := sarama.NewSyncProducer([]string{kafka.HOST}, config)
	if err != nil {
		log.Fatal("NewSyncProducer err:", err)
	}
	defer producer.Close()
	for i := 0; i < limit; i++ {
		msg := &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(defaultMsg)}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Println("SendMessage err: ", err)
			return
		}
		log.Printf("[Producer] partitionid: %d; offset:%d\n", partition, offset)
	}
}
