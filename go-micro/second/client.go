package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	pb "i-go/go-micro/second/pb"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "illusory"
)

func main() {
	// 我这里用的etcd 做为服务发现
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://192.168.0.2:32779", "http://192.168.0.2:32775", "http://192.168.0.2:32771",
		}
	})

	// 初始化服务
	// 	"github.com/micro/go-micro/service/grpc"
	// Go-grpc服务与go-micro服务一样，也就是说你可以直接将服务声明方式`micro.NewService`换成`grpc.NewService`，而不需要改动其它代码。
	service := grpc.NewService(
		micro.Name("go.micro.srv.hello"),
		// 注册服务的过期时间
		micro.RegisterTTL(time.Second*30),
		// 间隔多久再次注册服
		micro.RegisterInterval(time.Second*20),
		micro.Registry(reg),
	)

	service.Init()

	// Create new greeter client
	client := pb.NewHelloService("go.micro.srv.hello", service.Client())

	// Call the greeter
	for i := 0; i < 10; i++ {
		go func() {
			ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
			rsp, err := client.SayHello(ctx, &pb.User{Name: defaultName})
			if err != nil {
				logrus.Error(err)
			}
			// Print response
			fmt.Println(rsp)
		}()
	}
	fmt.Println("send req 10")
	time.Sleep(time.Second * 10)
}
