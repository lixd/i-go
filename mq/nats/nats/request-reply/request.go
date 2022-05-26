package request_reply

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"i-go/mq/nats/nats/conn"
)

var (
	nc  *nats.Conn
	err error
)

func init() {
	nc, err = conn.NewConn()
	if err != nil {
		panic(err)
	}
}

func Request(subject string, msg []byte) {
	resp, err := nc.Request(subject, msg, time.Second)
	if err != nil {
		logrus.WithField("scene", "nats publish").Error(err)
	}
	fmt.Printf("request %v reply:%+v\n", string(msg), string(resp.Data))
}
