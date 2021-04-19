package stan_sub

import (
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"i-go/mq/nats/constant"
	"i-go/mq/nats/stan/msghandler"
)

func TestStartMany(t *testing.T) {
	StartMany(2, constant.DefaultSubject, constant.DefaultQueue, msghandler.Simple)
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s
	Unsubscribe()
	time.Sleep(time.Second * 3)
	// 这里释放资源关闭连接什么的
}
