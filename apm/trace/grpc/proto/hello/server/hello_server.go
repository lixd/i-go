package main

import (
	"context"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"i-go/apm/trace/config"
	"i-go/apm/trace/grpc/interceptor"
	"i-go/grpc/hello/proto"
	"log"
	"net"
)

type helloServer struct{}

func (s *helloServer) SayHello(ctx context.Context, in *proto.HelloReq) (*proto.HelloRep, error) {
	return &proto.HelloRep{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", "50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	tracer, closer := config.NewTracer("gRPC-hello")
	defer closer.Close()
	// UnaryInterceptor
	s := grpc.NewServer(grpc.UnaryInterceptor(
		grpcMiddleware.ChainUnaryServer(
			interceptor.ServerInterceptor(tracer),
		),
	))
	proto.RegisterHelloServer(s, &helloServer{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
