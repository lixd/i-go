package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "i-go/grpc/proto"
	"log"
	"net"
)

const (
	port = ":50052"
)

type userServer struct {
}

func (s *userServer) Create(ctx context.Context, user *pb.User) (msg *pb.Resp, err error) {
	return &pb.Resp{Message: "Create Success"}, nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("net.Listen fail: %v", err)
	}
	newServer := grpc.NewServer()
	pb.RegisterUserServiceServer(newServer, &userServer{})
	newServer.Serve(listener)
}
