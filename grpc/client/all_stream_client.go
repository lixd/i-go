package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pro "i-go/grpc/proto"
	"i-go/utils"
	"io"
	"strconv"
	"sync"
	"time"
)

const AllStreamAddress = "localhost:50055"

/*
// 1. 建立连接 获取client
// 2. 调用方法获取stream
// 3. 开两个goroutine 分别用于Recv()和Send()
// 3.1 一直Recv()到err==io.EOF(即服务端关闭stream)
// 3.2 Send()则由自己控制
// 4. 发送完毕调用 stream.CloseSend()关闭stream 必须调用关闭 否则Server会一直尝试接收数据 一直报错...

*/
func main() {
	// 1.建立连接
	conn, err := grpc.Dial(AllStreamAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 1.new client
	client := pro.NewAllStreamServerClient(conn)
	waitGroup := sync.WaitGroup{}
	// 2. 调用方法获取stream
	stream, err := client.AllStream(context.Background())
	if err != nil {
		panic(err)
	}
	// 3.开两个goroutine 分别用于Recv()和Send()
	waitGroup.Add(1)
	go func() {
		for {
			data, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					fmt.Println("Server Closed")
					break
				}
				logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "AllStream Recv error"}).Error(err)
				continue
			}
			fmt.Printf("Recv Data:%v \n", data.Data)
		}
		waitGroup.Done()
	}()

	waitGroup.Add(1)
	go func() {
		for i := 0; i < 5; i++ {
			err := stream.Send(&pro.AllStreamReq{Data: strconv.Itoa(i)})
			if err != nil {
				logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "ClientStream Recv error"}).Error(err)
			}
			time.Sleep(time.Second)
		}
		// 4. 发送完毕关闭stream
		err := stream.CloseSend()
		if err != nil {
			logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "ClientStream CloseSend error"}).Error(err)
		}
		waitGroup.Done()
	}()
	waitGroup.Wait()
}
