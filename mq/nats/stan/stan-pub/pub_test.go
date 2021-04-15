package stan_pub

import (
	"testing"

	"i-go/mq/nats/constant"
)

func TestPublishMsg(t *testing.T) {
	for i := 0; i < 999; i++ {
		PublishMsg(constant.DefaultSubject, []byte("hello nats-streaming"))
	}
	Release()
}
