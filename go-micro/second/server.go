package main

import (
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
	port = ":50051"
)

// 定义一个结构体
type helloServerNew struct{}

// 实现proto文件中定义的方法
func (s *helloServerNew) SayHello(ctx context.Context, req *pb.User, resp *pb.Msg) error {
	resp.Message = "Hello " + req.Name
	return nil
}

func main() {
	// 我这里用的etcd 做为服务发现
	// github.com/micro/go-micro/registry
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
		micro.Version("latest"),
		micro.Registry(reg),
	)

	// Init方法会解析命令行flags
	service.Init(
		micro.AfterStart(func() error {
			logrus.Info("server img start...")
			return nil
		}),
		micro.AfterStop(func() error {
			logrus.Info("server img stop...")
			return nil
		}))

	err := pb.RegisterHelloHandler(service.Server(), &helloServerNew{})
	if err != nil {
		logrus.Error(err)
	}
	if err := service.Run(); err != nil {
		logrus.Error(err)
	}

	/*
		// 这里没必要了 内部已经做了这些工作
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		<-signalChan
		services, err := reg.GetService("go.micro.srv.hello")
		if err == nil && len(services) != 0 {
			s := services[0]
			_ = reg.Deregister(s)
		}*/
	// 实在不放心在加个延时即可
	time.Sleep(time.Second * 5)
}
