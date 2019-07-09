package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pro "i-go/grpc/proto"
	"net"
	"sync"
	"time"
)

const AllStreamPort = ":50055"

type allStream struct {
}

var AllStream = &allStream{}

// waitGroup 等待goroutine退出
func (server *allStream) AllStream(allStream pro.AllStreamServer_AllStreamServer) error {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	go func() {
		for {
			data, _ := allStream.Recv()
			fmt.Println(data)
		}
		waitGroup.Done()
	}()

	go func() {
		for {
			err := allStream.Send(&pro.AllStreamResp{Data: "All Stream From Server"})
			if err != nil {
				log.Error(err.Error())
			}
			time.Sleep(time.Second)
		}
		waitGroup.Done()
	}()
	// 让主线程卡在这里 如果两个 goroutine 都结束了 则结束主线程
	waitGroup.Wait()

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
