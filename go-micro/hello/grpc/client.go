package main

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"i-go/go-micro/hello/grpc/proto"
)

const (
	address = "localhost:50051"
)

func main() {
	// 开启一个链接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 用conn new一个client
	c := proto.NewGreeterClient(conn)
	// 用client 调用方法
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "lixd"})
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "grpc error"}).Error(err)
	}
	logrus.Infof("Greeting: %s", r.Message)
}
