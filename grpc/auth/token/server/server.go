package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"i-go/grpc/auth/token"
	pb "i-go/grpc/auth/token/proto"
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
	auth *token.Authentication
}

func (g *greeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// 增加身份校验(借助拦截器实现全局身份验证)
	if err := g.auth.Auth(ctx); err != nil {
		return nil, err
	}
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8087")
	if err != nil {
		panic(err)
	}
	newServer := grpc.NewServer()
	// 自定义验证新
	us := greeterServer{auth: &token.Authentication{Username: "17x", Password: "golang"}}
	pb.RegisterGreeterServer(newServer, &us)
	log.Println("Serving gRPC on 0.0.0.0:8087")
	if err = newServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
