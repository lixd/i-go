package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "i-go/grpc/proto"
	"net"
)

type userServer struct {
}

func (s *userServer) Create(ctx context.Context, req *pb.UserReq) (msg *pb.UserResp, err error) {
	fmt.Printf("Recv Age: %v Name:%v \n", req.Age, req.Name)
	return &pb.UserResp{Message: "Create Success"}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(err)
	}
	newServer := grpc.NewServer()
	pb.RegisterUserServiceServer(newServer, &userServer{})
	err = newServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
