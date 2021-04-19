package msghandler

import "github.com/sirupsen/logrus"

// Simple 简单的消息处理方法 可以看成消费者
func Simple(msg []byte) {
	logrus.Println(string(msg))
}
