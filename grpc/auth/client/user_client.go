package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"i-go/grpc/auth"
	"i-go/grpc/tls/proto"
	"i-go/utils"
)

/*
 */
func main() {
	credential := auth.Authentication{
		User:     "17x",
		Password: "golang",
	}
	// grpc.Dial 创建连接 grpc.WithBlock() 阻塞直到连接成功 WithTransportCredentials() tls
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithPerRPCCredentials(&credential))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewUserServiceClient(conn)
	resp, err := client.Create(context.Background(), &proto.UserReq{Name: "17x", Age: "23"})
	if err != nil {
		logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "create user  error"}).Error(err)
	}
	fmt.Printf("Create User Resp: %v \n", resp.Message)
}
