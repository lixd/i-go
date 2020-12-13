package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"i-go/grpc/auth"
	"i-go/grpc/tls/proto"
	"net"
)

type userServer struct {
	auth *auth.Authentication
}

func (s *userServer) Create(ctx context.Context, req *proto.UserReq) (msg *proto.UserResp, err error) {
	if err := s.auth.Auth(ctx); err != nil {
		return nil, err
	}
	fmt.Printf("Recv Age: %v Name:%v \n", req.Age, req.Name)
	return &proto.UserResp{Message: "Create Success"}, nil
}

/*
auth demo
*/
func main() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(err)
	}
	newServer := grpc.NewServer()
	us := userServer{auth: &auth.Authentication{User: "17x", Password: "golang"}}
	proto.RegisterUserServiceServer(newServer, &us)
	err = newServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
