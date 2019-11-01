package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	pb "i-go/go-micro/first/new_micro/pb"
)

const (
	address     = "localhost:50051"
	defaultName = "illusory"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("go.micro.srv.hello"))
	service.Init()

	// Create new greeter client
	client := pb.NewHelloService("go.micro.srv.hello", service.Client())

	// Call the greeter
	rsp, err := client.SayHello(context.TODO(), &pb.User{Name: defaultName})
	if err != nil {
		logrus.Error(err)
	}
	// Print response
	fmt.Println(rsp)
}
