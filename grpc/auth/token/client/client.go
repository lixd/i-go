package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"i-go/grpc/auth/token"
	pb "i-go/grpc/auth/token/proto"
)

func main() {
	credential := token.Authentication{
		Username: "17x1",
		Password: "golang",
	}
	//  WithTransportCredentials()  自定义验证
	conn, err := grpc.Dial("0.0.0.0:8087", grpc.WithInsecure(), grpc.WithPerRPCCredentials(&credential))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("SayHello error:%v", err)
	}
	log.Printf("Greeter: %v \n", resp.Message)
}
