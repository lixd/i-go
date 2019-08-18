package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	pb "i-go/go-micro/second/pb"
)

const (
	address     = "localhost:50051"
	defaultName = "illusory"
)

func main() {
	// 我这里用的etcd 做为服务发现，如果使用consul可以去掉
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://192.168.3.34:2379", "http://192.168.3.18:2379", "http://192.168.3.110:2379",
		}
	})

	// 初始化服务
	service := micro.NewService(
		micro.Registry(reg),
	)

	service.Init()

	// Create new greeter client
	client := pb.NewHelloService("go.micro.srv.hello", service.Client())

	// Call the greeter
	rsp, err := client.SayHello(context.TODO(), &pb.User{Name: defaultName})
	if err != nil {
		logrus.Error(err)
	}
	// Print response
	fmt.Println(rsp)
}
