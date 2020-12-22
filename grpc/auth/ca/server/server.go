package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "i-go/grpc/auth/ca/proto"
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (g *greeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	certificate, err := tls.LoadX509KeyPair("../cert/server.crt", "../cert/server.key")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../cert/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append certs")
	}
	// 构建基于 TLS 的 TransportCredentials
	creds := credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{certificate},
		// 要求必须校验客户端的证书 可以根据实际情况选用其他参数
		ClientAuth: tls.RequireAndVerifyClientCert, // NOTE: this is optional!
		// 设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
		ClientCAs: certPool,
	})
	server := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterGreeterServer(server, &greeterServer{})

	listener, err := net.Listen("tcp", ":8086")
	if err != nil {
		log.Fatalf("Listen err: %v", err)
	}
	log.Println("Serving gRPC on 0.0.0.0:8086")
	if err = server.Serve(listener); err != nil {
		panic(err)
	}
}
