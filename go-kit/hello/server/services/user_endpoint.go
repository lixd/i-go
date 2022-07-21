package services

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type UserRequest struct {
	UserId int `json:"userId"`
}
type UserResponse struct {
	Username string `json:"username"`
}

func GenUserEndpoint(service IUser) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		name := service.GetName(r.UserId)
		return UserResponse{Username: name}, nil
	}
}
