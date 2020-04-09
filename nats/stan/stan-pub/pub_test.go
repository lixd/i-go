package stan_pub

import (
	"i-go/nats/constant"
	"testing"
)

func TestPublishMsg(t *testing.T) {
	PublishMsg(constant.DefaultSubject, []byte("test msg"))
}
