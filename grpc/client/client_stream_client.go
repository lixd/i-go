package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pro "i-go/grpc/proto"
	"strconv"
	"time"
)

const ClientStreamAddress = "localhost:50054"

func main() {
	conn, err := grpc.Dial(ClientStreamAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer conn.Close()
	// 通过conn new 一个 client
	client := pro.NewClientStreamServerClient(conn)

	// 这里获取到的 res 是一个stream
	res, err := client.ClientStream(context.Background())
	if err != nil {
		log.Error(err.Error())
		return
	}

	for i := 0; i < 20; i++ {
		itoa := strconv.Itoa(i)
		// 通过 send 方法不断推送数据到server
		err := res.Send(&pro.ClientStreamReq{Data: " client data: " + itoa})
		if err != nil {
			log.Error(err.Error())
			return
		}
		time.Sleep(time.Second)
	}
	return
}
