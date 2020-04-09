package conn

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"i-go/nats/constant"
	"time"
)

func NewConn() (*nats.Conn, error) {
	totalWait := 3 * time.Minute
	reconnectDelay := time.Second
	nc, err := nats.Connect(
		StanURL(),
		nats.ReconnectWait(reconnectDelay),
		nats.MaxReconnects(int(totalWait/reconnectDelay)),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			logrus.Printf("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			logrus.Printf("Reconnected [%s]", nc.ConnectedUrl())
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			logrus.Fatalf("Exiting: %v", nc.LastError())
		}),
	)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "nats conn error"}).Error(err)
		return nil, err
	}
	return nc, nil
}

func StanURL() string {
	url := viper.GetString("stanurl")
	if url == "" {
		url = constant.DefaultNatsURL
	}
	return url
}
