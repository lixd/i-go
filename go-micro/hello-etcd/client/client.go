package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/sirupsen/logrus"
	"i-go/go-micro/hello-etcd/proto"
	"time"
)

const (
	Hello = "go.micro.srv.greeter"
)

func main() {
	// etcd服务注册与发现
	reg := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"123.57.236.125:12379", "123.57.236.125:22379", "123.57.236.125:32379"}
		options.Timeout = 10 * time.Second
	})
	service := micro.NewService(
		micro.Name(Hello),
		micro.Version("v1"),
		micro.Registry(reg),
	)
	service.Init()

	client := proto.NewGreeterService(Hello, service.Client())

	rsp, err := client.SayHello(context.TODO(), &proto.HelloRequest{Name: "lixd"})
	if err != nil {
		logrus.Error(err)
	}
	fmt.Println(rsp)
}
