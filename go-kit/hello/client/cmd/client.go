package main

import (
	"context"
	"fmt"
	httpTransport "github.com/go-kit/kit/transport/http"
	"i-go/go-kit/hello/client/services"
	"net/url"
	"os"
)

func main() {
	target, _ := url.Parse("http://localhost:8080")
	client := httpTransport.NewClient("GET", target, services.GetUserInfoReq, services.GetUserInfoResp)
	endpoint := client.Endpoint()
	i, err := endpoint(context.Background(), services.UserRequest{UserId: 999})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// response := i.(services.UserResponse)
	fmt.Println(i)
}
