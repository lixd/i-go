package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	pb "i-go/grpc/helloworld/proto"
)

func main() {
	conn, err := grpc.DialContext(context.Background(), "0.0.0.0:8080", grpc.WithInsecure(), grpc.WithDefaultServiceConfig())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", resp.Message)
}
