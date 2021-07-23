package async

import (
	"fmt"
	"log"
	"sync/atomic"

	"github.com/Shopify/sarama"
	"i-go/mq/kafka"
)

/*
	本例展示最简单的 异步生产者 的使用（除异步生产者外 kafka 还有同步生产者）
	名词 async producer
*/

var count int64

func Producer(topic string, limit int) {
	config := sarama.NewConfig()
	// 异步生产者不建议把 Errors 和 Successes 都开启，一般开启 Errors 就行
	// 同步生产者就必须都开启，因为会同步返回发送成功或者失败
	config.Producer.Return.Errors = true    // 设定需要返回错误信息
	config.Producer.Return.Successes = true // 设定需要返回成功信息
	// config.Producer.MaxMessageBytes = 1     // 设置消息体最大大小为 1 byte，用以发生错误
	producer, err := sarama.NewAsyncProducer([]string{kafka.HOST}, config)
	if err != nil {
		log.Fatal("NewSyncProducer err:", err)
	}
	defer producer.AsyncClose()
	// 异步发送
	for i := 0; i < limit; i++ {
		// str := strconv.Itoa(int(time.Now().UnixNano()))
		str := ""
		msg := &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(str)}
		// 异步发送只是写入内存了就返回了，并没有真正发送出去
		// sarama 库中用的是一个 channel 来接收，后台 goroutine 异步从该 channel 中取出消息并真正发送
		producer.Input() <- msg
		atomic.AddInt64(&count, 1)
		if atomic.LoadInt64(&count)%1000 == 0 {
			log.Printf("已发送消息数:%v\n", count)
		}
		// [!important] 异步生产者发送后必须把返回值从 Errors 或者 Successes 中读出来 不然会阻塞 producer.Input()
		// 成功的先不管，这里只打印错误新
		select {
		case _ = <-producer.Successes():
		case e := <-producer.Errors():
			if e != nil {
				fmt.Printf("[Producer] err:%v msg:%+v \n", e.Msg, e.Err)
			}
		}
	}
	log.Printf("发送完毕 总发送消息数:%v\n", limit)
}
