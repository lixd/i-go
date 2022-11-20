package conn

import (
	"time"

	"i-go/mq/nats/constant"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewConn() (*nats.Conn, error) {
	totalWait := 3 * time.Minute
	reconnectDelay := time.Second
	nc, err := nats.Connect(
		StanURL(),
		nats.UserInfo("admin", "think"),
		nats.ClientCert("../../../../dist/cert/client.crt", "../../../../dist/cert/client.key"),
		nats.RootCAs("../../../../dist/cert/ca.crt"),
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
	// return "nats://127.0.0.1:9889"
	return "nats://localhost:4222"
	// return "nats://172.20.149.197:4222"
	url := viper.GetString("stanurl")
	if url == "" {
		url = constant.DefaultNatsURL
	}
	return url
}
