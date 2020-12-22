package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "i-go/grpc/auth/tls/proto"
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (g *greeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

/*
1. credentials.NewServerTLSFromFile("../../conf/server.pem", "../../conf/server.key") 构建TransportCredentials
2. grpc.NewServer(grpc.Creds(c)) 开启TLS
*/
func main() {
	// 构建 TransportCredentials
	c, err := credentials.NewServerTLSFromFile("../cert/server.crt", "../cert/server.key")
	if err != nil {
		log.Fatalf("NewServerTLSFromFile err: %v", err)
	}
	listener, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatalf("Listen err: %v", err)
	}
	// 通过 grpc.Creds(c) 开启TLS
	newServer := grpc.NewServer(grpc.Creds(c))
	pb.RegisterGreeterServer(newServer, &greeterServer{})
	log.Println("Serving gRPC on 0.0.0.0:8085")
	if err = newServer.Serve(listener); err != nil {
		panic(err)
	}
}
