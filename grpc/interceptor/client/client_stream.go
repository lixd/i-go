package main

import (
	"io"
	"log"
	"sync"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"i-go/grpc/interceptor/inter"
	pb "i-go/grpc/interceptor/proto"
)

func main() {
	var (
		wg sync.WaitGroup
	)
	// 指定拦截器 这里指调用stream方法SayHello所以只需要注册stream相关拦截器
	streamInts := grpc.WithChainStreamInterceptor(inter.StreamClientRecovery, inter.StreamClientFilter, inter.StreamClientLogging)
	conn, err := grpc.DialContext(context.Background(), "0.0.0.0:8084", grpc.WithInsecure(),
		grpc.WithBlock(), streamInts)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewInterceptorClient(conn)
	stream, err := client.Sqrt(context.Background())
	if err != nil {
		log.Fatalf("Sqrt error: %v", err)
	}
	// 3.开两个goroutine 分别用于Recv()和Send()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			data, err := stream.Recv()
			if err == io.EOF {
				log.Println("Server Closed")
				break
			}
			if err != nil {
				continue
			}
			log.Printf("Recv Data:%v \n", data.Sqrt)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			err := stream.Send(&pb.SqrtReq{Number: float64(i)})
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
