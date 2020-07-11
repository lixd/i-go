package main

import (
	httpTransport "github.com/go-kit/kit/transport/http"
	"i-go/go-kit/hello/server/services"
	"net/http"
)

func main() {
	user := services.UserServer{}
	endpoint := services.GenUserEndpoint(&user)

	server := httpTransport.NewServer(endpoint, services.DecodeUserRequest, services.EncodeUserResponse)
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		panic(err)
	}
}
