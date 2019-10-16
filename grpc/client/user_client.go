package main

import (
	"context"
	"google.golang.org/grpc"
	pb "i-go/grpc/proto"
	"log"
)

func main() {
	// grpc.Dial 创建连接 grpc.WithInsecure() 禁用传输安全性
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)
	resp, err := client.Create(context.Background(), &pb.UserReq{Name: "illusory", Age: "23"})
	if err != nil {
		log.Fatalf("could not Create: %v", err)
	}
	log.Printf("Create Resp: %s", resp.Message)
}
