package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pro "i-go/grpc/proto"
	"net"
)

const ClientStreamPort = ":50054"

type clientStream struct {
}

var ClientStream = &clientStream{}

func (server *clientStream) ClientStream(res pro.ClientStreamServer_ClientStreamServer) error {
	for {
		// 通过 Recv() 不断获取服务端send()推送的消息
		data, err := res.Recv()
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println(data)
	}
}

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", ClientStreamPort)
	if err != nil {
		return
	}
	newServer := grpc.NewServer()
	// 注册server
	pro.RegisterClientStreamServerServer(newServer, &clientStream{})
	newServer.Serve(lis)
}
