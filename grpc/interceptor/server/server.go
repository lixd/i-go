package main

import (
	"io"
	"log"
	"math"
	"net"
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"i-go/grpc/interceptor/inter"
	pb "i-go/grpc/interceptor/proto"
)

type interceptor struct {
	pb.UnimplementedInterceptorServer
}

func (i *interceptor) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// 手动panic用于测试recovery拦截器
	// panic("test")
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
func (i *interceptor) Sqrt(stream pb.Interceptor_SqrtServer) error {
	// panic("test")

	var (
		waitGroup sync.WaitGroup
		numbers   = make(chan float64)
	)
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()

		for v := range numbers {
			err := stream.Send(&pb.SqrtReply{Sqrt: math.Sqrt(v)})
			if err != nil {
				log.Printf("Send error:%v \n", err)
				continue
			}
		}
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Recv error:%v", err)
			}
			log.Printf("Recv Data:%v \n", req.Number)
			numbers <- req.Number
		}
		close(numbers)
	}()
	waitGroup.Wait()

	// 返回nil表示已经完成响应
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8084")
	if err != nil {
		panic(err)
	}
	// https://github.com/grpc/grpc-go/pull/3336
	// gRPCv1.28.0 增加了ChainUnaryInterceptor 多Interceptor的情况也可以不借助 go-grpc-middleware 这个包了
	// interceptors := grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
	// 	inter.UnaryServerRecovery,inter.UnaryServerFilter,inter.UnaryServerLogging))
	// 服务端需要注册所有拦截器(unary和stream)
	// Recovery 必须放最前面才能捕获后续拦截器中触发的 panic
	unaryInts := grpc.ChainUnaryInterceptor(inter.UnaryServerRecovery, inter.UnaryServerFilter, inter.UnaryServerLogging)
	streamInts := grpc.ChainStreamInterceptor(inter.StreamServerRecovery, inter.StreamServerFilter, inter.StreamServerLogging)
	s := grpc.NewServer(unaryInts, streamInts)
	pb.RegisterInterceptorServer(s, &interceptor{})
	log.Println("Serving gRPC on 0.0.0.0:8084")
	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}
