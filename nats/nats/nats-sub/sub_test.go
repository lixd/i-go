package nats_sub

import (
	"i-go/nats/constant"
	"i-go/nats/nats/msghandler"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestSubscribe(t *testing.T) {
	Subscribe(constant.DefaultSubject, msghandler.Simple)
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s
	UnSubscribe()
}
