package main

import (
	"context"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "i-go/grpc-demo/helloworld"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
	parallel    = 20000
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	var wg sync.WaitGroup
	wg.Add(parallel)
	for i := 0; i < parallel; i++ {
		go call(c, &wg)
	}
	wg.Wait()
}

func call(c pb.GreeterClient, wg *sync.WaitGroup) {
	defer wg.Done()
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: defaultName})
	if err != nil {
		log.Printf("could not greet: %v", err)
	}
	_ = r
	log.Printf("Greeting: %s", r.GetMessage())
}
