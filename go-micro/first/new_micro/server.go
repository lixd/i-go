package main

import (
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	pb "i-go/go-micro/first/new_micro/pb"
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
	// 注意，在这里我们使用go-micro的NewService方法来创建新的微服务服务器，
	// 而不是上一篇文章中所用的标准
	srv := micro.NewService(
		// 注意，Name方法的必须是你在proto文件中定义的package名字
		micro.Name("go.micro.srv.hello"),
		micro.Version("latest"),
	)
	// Init方法会解析命令行flags
	srv.Init()
	err := pb.RegisterHelloHandler(srv.Server(), &helloServerNew{})
	if err != nil {
		logrus.Error(err)
	}
	if err := srv.Run(); err != nil {
		logrus.Error(err)
	}
}
