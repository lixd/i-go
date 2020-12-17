package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "i-go/grpc/stream/proto"
)

/*
1. 建立连接 获取client
2. 调用方法获取stream
3. 开两个goroutine 分别用于Recv()和Send()
	3.1 一直Recv()到err==io.EOF(即服务端关闭stream)
	3.2 Send()则由自己控制
4. 发送完毕调用 stream.CloseSend()关闭stream 必须调用关闭 否则Server会一直尝试接收数据 一直报错...

*/
func main() {
	var (
		wg sync.WaitGroup
	)

	// 1.建立连接
	conn, err := grpc.DialContext(context.Background(), "0.0.0.0:8083", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 1.new client
	client := pb.NewBidirectionalStreamServerClient(conn)
	// 2. 调用方法获取stream
	stream, err := client.Sqrt(context.Background())
	if err != nil {
		panic(err)
	}
	// 3.开两个goroutine 分别用于Recv()和Send()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			data, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("Server Closed")
				break
			}
			if err != nil {
				continue
			}
			fmt.Printf("Recv Data:%v \n", data.Sqrt)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			err := stream.Send(&pb.AllStreamReq{Number: float64(i)})
			if err != nil {
				log.Printf("Send error:%v\n", err)
			}
			time.Sleep(time.Second)
		}
		// 4. 发送完毕关闭stream
		err := stream.CloseSend()
		if err != nil {
			log.Printf("Send error:%v\n", err)
			return
		}
	}()
	wg.Wait()
}
