package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	pb "i-go/grpc/stream/proto"
)

/*
1. 建立连接并获取client
2. 获取 stream 并通过 Send 方法不断推送数据到服务端
3. 发送完成后通过stream.CloseAndRecv() 关闭steam并接收服务端返回结果
*/
func main() {
	// 1.建立连接并获取 client
	conn, err := grpc.DialContext(context.Background(), "0.0.0.0:8081", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewClientStreamClient(conn)

	// 2.获取 stream 并通过 Send 方法不断推送数据到服务端
	stream, err := client.Sum(context.Background())
	if err != nil {
		log.Fatalf("Sum() error: %v", err)
	}
	for i := int64(0); i < 10; i++ {
		err := stream.Send(&pb.ClientStreamReq{Number: i})
		if err != nil {
			log.Printf("Send(%v) error: %v", i, err)
			continue
		}
	}

	// 3. 发送完成后通过stream.CloseAndRecv() 关闭steam并接收服务端返回结果
	// (服务端则根据err==io.EOF来判断client是否关闭stream)
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("CloseAndRecv() error: %v", err)
	}
	log.Printf("sum: %v", resp.GetSum())
}
