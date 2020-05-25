package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/sirupsen/logrus"
	"i-go/go-micro/hello/micro/proto"
)

const (
	Hello = "go.micro.srv.greeter"
)

func main() {
	service := micro.NewService(micro.Name(Hello))
	service.Init()

	client := proto.NewGreeterService(Hello, service.Client())

	rsp, err := client.SayHello(context.TODO(), &proto.HelloRequest{Name: "lixd"})
	if err != nil {
		logrus.Error(err)
	}
	fmt.Println(rsp)
}
