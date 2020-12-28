package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "i-go/grpc/auth/ca/proto"
)

func main() {
	certificate, err := tls.LoadX509KeyPair("../cert/client.crt", "../cert/client.key")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../cert/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   "www.lixueduan.com", // NOTE: this is required!
		RootCAs:      certPool,
	})
	conn, err := grpc.DialContext(context.Background(), "0.0.0.0:8086", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("DialContext error:%v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("SayHello error:%v", err)
	}
	log.Printf("Greeter: %v \n", resp.Message)
}
