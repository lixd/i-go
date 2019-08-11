package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "i-go/go-micro/first/old_grpc/pb"
	"log"
)

const (
	address     = "localhost:50051"
	defaultName = "illusory"
)

func main() {
	// 开启一个链接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// 用conn new一个client
	c := pb.NewHelloClient(conn)
	// 用client 调用方法
	r, err := c.SayHello(context.Background(), &pb.User{Name: defaultName})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
