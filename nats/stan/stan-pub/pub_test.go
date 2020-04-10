package stan_pub

import (
	"i-go/nats/constant"
	"testing"
)

func TestPublishMsg(t *testing.T) {
	for i := 0; i < 999; i++ {
		PublishMsg(constant.DefaultSubject, []byte("hello nats-streaming"))
	}
	Release()
}
