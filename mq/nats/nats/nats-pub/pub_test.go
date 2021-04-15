package nats_pub

import (
	"testing"

	"i-go/mq/nats/constant"
)

func TestPublishMsg(t *testing.T) {
	for i := 0; i < 99; i++ {
		PublishMsg(constant.DefaultSubject, []byte("hello nats"))
	}
	Release()
}
