package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"i-go/go-micro/hello/grpc/proto"
	"net"
)

const (
	port = ":50051"
)

// 定义一个结构体
type helloServer struct{}

// 实现proto文件中定义的方法
func (s *helloServer) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	// 监听
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()

	// 在我们的gRPC服务器上注册微服务，这会将我们的代码和*.pb.go中
	// 的各种interface对应起来
	// 和http的路由差不多 用一个map结构来存储的 key就是服务名字 value是内部拼装起来的server
	proto.RegisterGreeterServer(s, &helloServer{})

	// 在gRPC服务器上注册reflection 这个方法和上面的pb.RegisterHelloServer逻辑一样的 内部调用的其实是一个方法..
	// 不过这里是注册的内部的一个默认的服务 &serverReflectionServer{ s: s}
	// 如果启动了gprc反射服务,那么就可以通过reflection包提供的反射服务查询GRPC服务或调用GRPC方法。
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
	}
}
