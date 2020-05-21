package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"i-go/grpc/hello/proto"
	"net"
	"runtime/debug"
)

const (
	port = ":50051"
)

// gPRC中的interceptor
// 定义一个结构体
type helloServer struct{}

// 实现proto文件中定义的方法
func (s *helloServer) SayHello(ctx context.Context, in *proto.HelloReq) (*proto.HelloRep, error) {
	//
	if ctx.Err() == context.Canceled {
		return &proto.HelloRep{}, nil
	}
	return &proto.HelloRep{Message: "Hello " + in.Name}, nil
}

func main() {
	// RPC
	// 监听
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	// 构建Server的时候传入两个写好的interceptor
	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(LoggingInterceptor, RecoveryInterceptor)))
	// 注册 server
	proto.RegisterHelloServer(s, &helloServer{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}

// LoggingInterceptor RPC 方法的入参出参的日志输出
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logrus.Printf("gRPC before: %s, %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	logrus.Printf("gRPC after: %s, %v", info.FullMethod, resp)
	return resp, err
}

// RecoveryInterceptor RPC 方法的异常保护和日志输出
func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "Panic err: %v", e)
		}
	}()

	return handler(ctx, req)
}
