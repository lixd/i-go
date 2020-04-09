package msghandler

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

// Simple 简单的消息处理方法 可以看成消费者
func Simple(msg *nats.Msg) {
	logrus.Println(string(msg.Data))
}
