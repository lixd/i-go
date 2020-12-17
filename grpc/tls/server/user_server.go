package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"i-go/grpc/tls"
	"i-go/grpc/tls/proto"
	"net"
)

type userServer struct {
}

func (s *userServer) Create(ctx context.Context, req *proto.UserReq) (msg *proto.UserResp, err error) {
	fmt.Printf("Recv Age: %v Name:%v \n", req.Age, req.Name)
	return &proto.UserResp{Message: "Create Success"}, nil
}

/*
1. credentials.NewServerTLSFromFile("../../conf/server.pem", "../../conf/server.key") 构建TransportCredentials
2. grpc.NewServer(grpc.Creds(c)) 开启tls
*/
func main() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(err)
	}
	crt, key := tls.Server()
	// 构建TransportCredentials
	//c, err := credentials.NewServerTLSFromFile("../../conf/server.pem", "../../conf/server.key")
	c, err := credentials.NewServerTLSFromFile(crt, key)
	if err != nil {
		logrus.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
	}
	// 开启tls
	newServer := grpc.NewServer(grpc.Creds(c))
	// newServer := grpc.NewServer()
	proto.RegisterUserServiceServer(newServer, &userServer{})
	err = newServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
