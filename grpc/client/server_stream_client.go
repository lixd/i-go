package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pro "i-go/grpc/proto"
	"time"
)

const ServerStreamAddress = "localhost:50053"

func main() {
	conn, err := grpc.Dial(ServerStreamAddress, grpc.WithInsecure())
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer conn.Close()
	// 通过conn new 一个 client
	client := pro.NewMyStreamServerClient(conn)

	data := &pro.ServerStreamReq{Data: "1"}
	// 这里获取到的 res 是一个stream
	res, err := client.ServerStream(context.Background(), data)
	if err != nil {
		log.Error(err.Error())
		return
	}

	start := time.Now().Unix()
	for {
		// 通过 Recv() 不断获取服务端send()推送的消息
		data, err := res.Recv()
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println(data)
	}
	fmt.Printf("take time :%v", time.Now().Unix()-start)
}
