package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pb "i-go/grpc/proto"
	"i-go/utils"
)

func main() {
	// grpc.Dial 创建连接 grpc.WithInsecure() 禁用传输安全性
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)
	resp, err := client.Create(context.Background(), &pb.UserReq{Name: "illusory", Age: "23"})
	if err != nil {
		logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "create user  error"}).Error(err)
	}
	fmt.Printf("Create User Resp: %v \n", resp.Message)
}
