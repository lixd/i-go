package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/sirupsen/logrus"
	"i-go/go-micro/hello/micro/proto"
)

const (
	Hello = "go.micro.srv.greeter"
)

// 定义一个结构体
type helloServerNew struct{}

// 实现proto文件中定义的方法
func (s *helloServerNew) SayHello(ctx context.Context, in *proto.HelloRequest, out *proto.HelloReply) error {
	out.Message = "Hello " + in.Name
	return nil
}

/*
使用的micro内部的注册中心
*/
func main() {
	// 注意，在这里我们使用go-micro的NewService方法来创建新的微服务服务器，
	// 而不是上一篇文章中所用的标准
	srv := micro.NewService(
		// 名字需要注意 微服务是通过名字调用的 所以客户端调用时也要用同样的名字
		micro.Name(Hello),
		micro.Version("latest"),
	)

	/*
		//当然这里也可以直接用grpc.NewServer()
		//不过这个grpc 是"github.com/micro/go-micro/service/grpc"
		// gomicro为了兼容grpc做的特殊处理 内部也是调用的micro.NewService()
		srv := grpc.NewService(
			micro.Name("go.micro.srv.hello"),
			micro.Version("latest"),
		)*/
	// Init方法会解析命令行flags
	srv.Init()
	// 这里是注册 可以看成http server的注册路由
	// 实际内部执行的是传入的第一个参数 即这里的srv.Server() 的Handle(Handler) error方法
	// gomicro这里有3个实现 rpc grpc和一个mock 应该用来测试的吧
	err := proto.RegisterGreeterHandler(srv.Server(), &helloServerNew{})
	if err != nil {
		logrus.Error(err)
	}
	// run方法调用了start方法
	// start方法也和上面一样有3个实现 grpc的start方法内部和普通grpc差不多
	// 也是net.Listen开启监听	ts, err := net.Listen("tcp", config.Address)
	// 然后还调用了register方法 根据不同的注册中心也有不同实现 就是把server信息添加到注册中心去
	// 然后循环根据设置的注册时间间隔 每过一段时间就去重新注册一次
	// start方法中同时也监听了 linux下的syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT这几个信号
	// 信号和context.done触发时都会执行stop方法 stop方法也是3种实现 grpc的stop方法是把started状态改成false然后发送了一个exit信号
	// run方法循环注册那个地方接收到这个退出信号也就退出循环了 然后从注册中心取消注册 接着有一个优雅关闭
	// 优雅关闭大概就是先不在接收新的连接 然后等所有的连接都处理完了再关掉
	if err := srv.Run(); err != nil {
		logrus.Error(err)
	}
}
