package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pro "i-go/grpc/proto"
	"i-go/utils"
	"io"
	"net"
	"strconv"
	"sync"
	"time"
)

const AllStreamPort = ":50055"

type allStream struct {
}

var AllStream = &allStream{}

// AllStream 双向流服务端
/*
// 1. 建立连接 获取client
// 2. 调用方法获取stream
// 3. 开两个goroutine 分别用于Recv()和Send()
// 3.1 一直Recv()到err==io.EOF(即客户端关闭stream)
// 3.2 Send()则自己控制什么时候Close 服务端stream没有close 只要跳出循环就算close了。 具体见https://github.com/grpc/grpc-go/issues/444
*/
func (server *allStream) AllStream(stream pro.AllStreamServer_AllStreamServer) error {
	waitGroup := sync.WaitGroup{}

	waitGroup.Add(1)
	go func() {
		for i := 0; i < 5; i++ {
			err := stream.Send(&pro.AllStreamResp{Data: strconv.Itoa(i)})
			if err != nil {
				logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "AllStream Send error"}).Error(err)
				continue
			}
			time.Sleep(time.Second)
		}
		waitGroup.Done()
	}()

	waitGroup.Add(1)
	go func() {
		for {
			data, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					fmt.Println("Client Closed")
					break
				}
				logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "AllStream Recv error"}).Error(err)
				continue
			}
			fmt.Printf("Recv Data:%v \n", data.Data)
		}
		waitGroup.Done()
	}()

	waitGroup.Wait()

	return nil

}

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", AllStreamPort)
	if err != nil {
		panic(err)
	}
	newServer := grpc.NewServer()
	// 注册server
	pro.RegisterAllStreamServerServer(newServer, &allStream{})
	err = newServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
