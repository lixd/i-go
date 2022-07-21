package main

import (
	"net/http"

	"i-go/go-kit/hello/server/services"

	httpTransport "github.com/go-kit/kit/transport/http"
)

func main() {
	user := services.UserServer{}
	endpoint := services.GenUserEndpoint(&user)
	server := httpTransport.NewServer(endpoint, services.DecodeUserRequest, services.EncodeUserResponse)
	if err := http.ListenAndServe(":8080", server); err != nil {
		panic(err)
	}
}
