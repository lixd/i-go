package main

import (
	"context"
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

type helloServerNew struct{}

func (s *helloServerNew) SayHello(ctx context.Context, in *proto.HelloRequest, out *proto.HelloReply) error {
	out.Message = "Hello " + in.Name
	return nil
}

/*
使用的etcd做注册中心
*/
func main() {
	// etcd-->"github.com/micro/go-micro/v2/registry/etcd"
	// NewRegistry 需要传一个或多个func进去 type Option func(*Options)
	reg := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"123.57.236.125:12379", "123.57.236.125:22379", "123.57.236.125:32379"}
		options.Timeout = 10 * time.Second
	})

	// 注意，在这里我们使用go-micro的NewService方法来创建新的微服务服务器，
	/*
		当然这里也可以直接用grpc.NewService()
		不过这个grpc 是"github.com/micro/go-micro/service/grpc"
		gomicro为了兼容grpc做的特殊处理 内部也是调用的micro.NewService()
		srv := grpc.NewService(
			micro.Name("go.micro.srv.hello"),
			micro.Version("latest"),
		)*/
	srv := micro.NewService(
		// 名字需要注意 客户端是通过名字调用的
		micro.Name(Hello),
		micro.Version("v1"),
		micro.Registry(reg),
		micro.AfterStart(func() error {
			logrus.WithFields(logrus.Fields{"Server": Hello, "Scenes": "server start..."})
			return nil
		}),
		micro.AfterStop(func() error {
			logrus.WithFields(logrus.Fields{"Server": Hello, "Scenes": "server stop..."})
			return nil
		}),
	)

	// Init方法会解析命令行flags
	srv.Init()
	// 这里是注册 可以看成http server的注册路由
	// 实际内部执行的是传入的第一个参数 即这里的srv.Server() 的Handle(Handler) error方法
	// gomicro这里有3个实现 rpc grpc和一个mock 应该用来测试的吧
	err := proto.RegisterGreeterHandler(srv.Server(), &helloServerNew{})
	if err != nil {
		panic(err)
	}

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
