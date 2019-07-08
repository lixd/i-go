package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "i-go/grpc/proto"
	"log"
	"net"
)

type userServer struct {
}

func (s *userServer) Create(ctx context.Context, user *pb.UserReq) (msg *pb.UserResp, err error) {
	return &pb.UserResp{Message: "Create Success"}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("net.Listen fail: %v", err)
	}
	newServer := grpc.NewServer()
	pb.RegisterUserServiceServer(newServer, &userServer{})
	newServer.Serve(listener)
}
