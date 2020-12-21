package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "i-go/grpc/helloworld/proto"
)

// greeterServer 随便定义一个结构体用于实现 .proto文件中定义的API
// 新版本 gRPC 要求必须嵌入 pb.UnimplementedGreeterServer 对象
type greeterServer struct {
	pb.UnimplementedGreeterServer
}

// SayHello 简单实现一下.proto文件中定义的 API
func (g *greeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	// 将服务描述(server)及其具体实现(greeterServer)注册到 gRPC 中去.
	// 内部使用的是一个 map 结构存储，类似 HTTP server。
	pb.RegisterGreeterServer(server, &greeterServer{})
	// 注册反射服务器 用于获取服务信息,和pb.RegisterGreeterServer()类似,只是没有具体实现
	reflection.Register(server)
	log.Println("Serving gRPC on 0.0.0.0:8080")
	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
