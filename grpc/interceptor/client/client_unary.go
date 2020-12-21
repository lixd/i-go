package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"i-go/grpc/interceptor/inter"
	pb "i-go/grpc/interceptor/proto"
)

func main() {
	// 指定拦截器 这里指调用unary方法SayHello所以只需要注册unary相关拦截器
	unaryInts := grpc.WithChainUnaryInterceptor(inter.UnaryClientRecovery, inter.UnaryClientFilter, inter.UnaryClientLogging)
	conn, err := grpc.DialContext(context.Background(), "0.0.0.0:8084", grpc.WithInsecure(),
		grpc.WithBlock(), unaryInts)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewInterceptorClient(conn)
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", resp.Message)
}
