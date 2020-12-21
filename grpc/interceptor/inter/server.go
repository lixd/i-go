package inter

import (
	"context"
	"log"
	"runtime/debug"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UnaryServerFilter 一元拦截器
/*
服务端一元拦截器需要实现为 grpc.UnaryServerInterceptor 这个类型 服务端流拦截器则是 grpc.StreamServerInterceptor 类型
ctx、req 则是 gPRC 方法的前两个参数
info参数包含了当前对应的那个gRPC方法各种信息
handler参数对应当前的gRPC方法函数
*/
func UnaryServerFilter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
	resp interface{}, err error) {
	log.Printf("unary filter server:%v method:%v :", info.Server, info.FullMethod)
	return handler(ctx, req)
}
func StreamServerFilter(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("stream filter method:%v :", info.FullMethod)
	return handler(srv, ss)
}

// UnaryServerLogging RPC 方法的入参出参的日志输出
func UnaryServerLogging(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
	interface{}, error) {
	log.Printf("unary gRPC before: %s, %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	log.Printf("unary gRPC after: %s, %v", info.FullMethod, resp)
	return resp, err
}
func StreamServerLogging(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("stream gRPC before: %s", info.FullMethod)
	err := handler(srv, ss)
	log.Printf("stream gRPC after: %s", info.FullMethod)
	return err
}

// UnaryServerRecovery RPC 方法的异常保护和日志输出
func UnaryServerRecovery(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
	resp interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "unary panic err: %v", e)
		}
	}()

	return handler(ctx, req)
}

func StreamServerRecovery(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	var err error
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "stream panic err: %v", e)
		}
	}()
	err = handler(srv, ss)
	return err
}
