package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "i-go/grpc/auth/tls/proto"
)

/*
1. credentials.NewClientTLSFromFile(crt, "www.lixueduan.com") 构建TransportCredentials
2. rpc.WithTransportCredentials(c) 配置TLS
*/
func main() {
	// serverNameOverride(即这里的"grpc")需要和生成证书时指定的Common Name对应
	c, err := credentials.NewClientTLSFromFile("../cert/server.crt", "www.lixueduan.com")
	// c, err := credentials.NewClientTLSFromFile(crt, "grpc")
	if err != nil {
		log.Fatalf("NewClientTLSFromFile error:%v", err)
	}
	// grpc.WithTransportCredentials(c) 配置TLS
	conn, err := grpc.DialContext(context.Background(), "0.0.0.0:8085", grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("DialContext error:%v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("SayHello error:%v", err)
	}
	log.Printf("Greeter: %v \n", resp.Message)
}
