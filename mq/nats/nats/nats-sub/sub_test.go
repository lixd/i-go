package nats_sub

import (
	"os"
	"os/signal"
	"syscall"
	"testing"

	"i-go/mq/nats/constant"
	"i-go/mq/nats/nats/msghandler"
)

func TestSubscribe(t *testing.T) {
	Subscribe(constant.DefaultSubject, msghandler.Simple)
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s
	UnSubscribe()
}
