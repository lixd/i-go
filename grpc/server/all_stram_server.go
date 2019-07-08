package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pro "i-go/grpc/proto"
	"net"
	"time"
)

const AllStreamPort = ":50055"

type allStream struct {
}

var AllStream = &allStream{}

func (server *allStream) AllStream(allStream pro.AllStreamServer_AllStreamServer) error {
	ok := make(chan bool, 2)
	go func() {
		for {
			data, _ := allStream.Recv()
			fmt.Println(data)
		}
		ok <- true
	}()

	go func() {
		for {
			err := allStream.Send(&pro.AllStreamResp{Data: "All Stream From Server"})
			if err != nil {
				log.Error(err.Error())
			}
			time.Sleep(time.Second)
		}
		ok <- true
	}()
	// 让主线程卡在这里 如果两个 goroutine 都结束了 则结束主线程
	for i := 0; i < 2; i++ {
		<-ok
	}
	return nil

}

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", AllStreamPort)
	if err != nil {
		return
	}
	newServer := grpc.NewServer()
	// 注册server
	pro.RegisterAllStreamServerServer(newServer, &allStream{})
	newServer.Serve(lis)
}
