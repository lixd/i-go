package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {
	// 1.连接到默认服务器
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	// json encoder
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	// Define the object
	type stock struct {
		Symbol string
		Price  int
	}

	// Publish the message
	if err := ec.Publish("updates", &stock{Symbol: "GOOG", Price: 1200}); err != nil {
		log.Fatal(err)
	}
	if err := nc.Publish("updates", []byte("All is Well")); err != nil {
		log.Fatal(err)
	}
	// 收件箱
	inbox := nats.NewInbox()
	subscription, err := nc.SubscribeSync(inbox)
	if err != nil {
		log.Fatal(err)
	}
	err = nc.PublishRequest("time", inbox, nil)
	if err != nil {
		log.Fatal(err)
	}
	msg, err := subscription.NextMsg(time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Reply: %s", msg.Data)

}
