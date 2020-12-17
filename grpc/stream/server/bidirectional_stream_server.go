package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"sync"

	"google.golang.org/grpc"
	pb "i-go/grpc/stream/proto"
)

type bidirectionalStream struct {
	pb.UnimplementedBidirectionalStreamServerServer
}

// Sqrt 双向流服务端
/*
// 1. 建立连接 获取client
// 2. 调用方法获取stream
// 3. 开两个goroutine（使用 chan 传递数据） 分别用于Recv()和Send()
// 3.1 一直Recv()到err==io.EOF(即客户端关闭stream)
// 3.2 Send()则自己控制什么时候Close 服务端stream没有close 只要跳出循环就算close了。 具体见https://github.com/grpc/grpc-go/issues/444
*/
func (b *bidirectionalStream) Sqrt(stream pb.BidirectionalStreamServer_SqrtServer) error {
	var (
		waitGroup sync.WaitGroup
		numbers   = make(chan float64)
	)
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()

		for v := range numbers {
			err := stream.Send(&pb.AllStreamResp{Sqrt: math.Sqrt(v)})
			if err != nil {
				fmt.Println("Send error:", err)
				continue
			}
		}
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Recv error:%v", err)
			}
			fmt.Printf("Recv Data:%v \n", req.Number)
			numbers <- req.Number
		}
		close(numbers)
	}()
	waitGroup.Wait()

	// 返回nil表示已经完成响应
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		panic(err)
	}
	newServer := grpc.NewServer()
	pb.RegisterBidirectionalStreamServerServer(newServer, &bidirectionalStream{})
	log.Println("Serving gRPC on 0.0.0.0:8083")
	if err = newServer.Serve(lis); err != nil {
		panic(err)
	}
}
