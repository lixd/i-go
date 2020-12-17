package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
	pb "i-go/grpc/stream/proto"
)

/*
1. 建立连接 获取client
2. 调用方法获取stream
3. for循环中通过stream.Recv()获取服务端推送的消息
4. err==io.EOF则表示服务端关闭stream了
*/
func main() {
	// 1.建立连接 获取client
	conn, err := grpc.DialContext(context.Background(), "0.0.0.0:8082", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewServerStreamClient(conn)
	// 2.调用获取stream
	stream, err := client.Pow(context.Background(), &pb.ServerStreamReq{Number: 2})
	if err != nil {
		log.Fatalf("Pow error:%v", err)
	}

	// 3. for循环获取服务端推送的消息
	for {
		// 3.通过 Recv() 不断获取服务端send()推送的消息
		resp, err := stream.Recv()
		// 4. err==io.EOF则表示服务端关闭stream了 退出
		if err == io.EOF {
			log.Println("server closed")
			break
		}
		if err != nil {
			log.Printf("Recv error:%v", err)
			continue
		}
		log.Printf("Recv data:%v", resp.Number)
	}
}
