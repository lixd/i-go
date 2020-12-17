package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "i-go/grpc/gateway/proto/helloworld"
)

const (
	address     = "localhost:8080"
	defaultName = "world"
)

func main() {
	conn, err := grpc.DialContext(context.Background(), address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: defaultName})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
