package main

import (
	"github.com/nats-io/nats.go"
	"log"
)

func main() {
	//1.连接到默认服务器
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
}
