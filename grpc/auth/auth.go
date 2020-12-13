package auth

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Authentication 需要实现 credentials.PerRPCCredentials 接口
type Authentication struct {
	User     string
	Password string
}

func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"user": a.User, "password": a.Password}, nil
}
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}

func (a *Authentication) Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}
	var (
		user     string
		password string
	)

	if val, ok := md["user"]; ok {
		user = val[0]
	}
	if val, ok := md["password"]; ok {
		password = val[0]
	}

	if user != a.User || password != a.Password {
		return status.Errorf(codes.Unauthenticated, "invalid token")
	}

	return nil
}
