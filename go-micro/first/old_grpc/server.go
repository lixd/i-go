package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "i-go/go-micro/first/old_grpc/pb"
	"log"
	"net"
)

const (
	port = ":50051"
)

// 定义一个结构体
type helloServer struct{}

// 实现proto文件中定义的方法
func (s *helloServer) SayHello(ctx context.Context, in *pb.User) (*pb.Msg, error) {
	return &pb.Msg{Message: "Hello " + in.Name}, nil
}

func main() {
	// 监听
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// 在我们的gRPC服务器上注册微服务，这会将我们的代码和*.pb.go中
	// 的各种interface对应起来
	pb.RegisterHelloServer(s, &helloServer{})

	// 在gRPC服务器上注册reflection
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
	}
}
