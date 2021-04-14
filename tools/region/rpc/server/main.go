package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"i-go/core/conf"
	"i-go/tools/region/core"
	pb "i-go/tools/region/rpc/proto"
)

const (
	port = ":50051"
)

// greeterServer 定义一个结构体用于实现 .proto文件中定义的方法
// 新版本 gRPC 要求必须嵌入 pb.UnimplementedGreeterServer 结构体
type regionServer struct {
	pb.UnimplementedRegionServerServer
}

// SayHello 简单实现一下.proto文件中定义的 SayHello 方法
func (r *regionServer) IP2Region(ctx context.Context, in *pb.IP) (*pb.Region, error) {
	log.Printf("Received: %v", in.GetIp())
	region := core.IP2Region(in.GetIp())
	return &pb.Region{Region: region}, nil
}

func main() {
	err := conf.Load("conf/config.yml")
	if err != nil {
		panic(err)
	}
	core.Init()

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// 将服务描述(server)及其具体实现(greeterServer)注册到 gRPC 中去.
	// 内部使用的是一个 map 结构存储，类似 HTTP server。
	pb.RegisterRegionServerServer(s, &regionServer{})
	log.Println("Serving gRPC on 0.0.0.0" + port)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
