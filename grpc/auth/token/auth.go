package token

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Authentication 需要实现 credentials.PerRPCCredentials 接口
type Authentication struct {
	Username string
	Password string
}

// GetRequestMetadata 获取授权信息并通过map返回出去，后续授权的时候使用
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"username": a.Username, "password": a.Password}, nil
}

// RequireTransportSecurity 是否需要开启TLS
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}

// Auth 具体的验证逻辑
func (a *Authentication) Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}
	var (
		user     string
		password string
	)

	if val, ok := md["username"]; ok {
		user = val[0]
	}
	if val, ok := md["password"]; ok {
		password = val[0]
	}

	if user != a.Username || password != a.Password {
		return status.Errorf(codes.Unauthenticated, "Unauthorized")
	}

	return nil
}
