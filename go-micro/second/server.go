package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	pb "i-go/go-micro/second/pb"
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
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://192.168.1.9:32772", "http://192.168.1.9:32773", "http://192.168.1.9:32769",
		}
	})

	// 初始化服务
	service := micro.NewService(
		micro.Name("go.micro.srv.hello"),
		micro.Version("latest"),
		micro.Registry(reg),
	)

	// Init方法会解析命令行flags
	service.Init()

	err := pb.RegisterHelloHandler(service.Server(), &helloServerNew{})
	if err != nil {
		logrus.Error(err)
	}
	if err := service.Run(); err != nil {
		logrus.Error(err)
	}
}
