package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"i-go/grpc/hello/proto"
	"log"
	"net"
)

const (
	port = ":50051"
)

// 定义一个结构体
type helloServer struct{}

// 实现proto文件中定义的方法
func (s *helloServer) SayHello(ctx context.Context, in *proto.HelloReq) (*proto.HelloRep, error) {
	return &proto.HelloRep{Message: "Hello " + in.Name}, nil
}

func main() {
	// RPC
	// 监听
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// 注册 server
	proto.RegisterHelloServer(s, &helloServer{})
	s.Serve(lis)

}
