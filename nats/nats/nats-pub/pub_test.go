package nats_pub

import (
	"i-go/nats/constant"
	"testing"
)

func TestPublishMsg(t *testing.T) {
	for i := 0; i < 99; i++ {
		PublishMsg(constant.DefaultSubject, []byte("hello nats"))
	}
	Release()
}
