package inter

import (
	"context"
	"log"
	"runtime/debug"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UnaryFilter 一元拦截器
/*
服务端一元拦截器需要实现为 grpc.UnaryServerInterceptor 这个类型 服务端流拦截器则是 grpc.StreamServerInterceptor 类型
ctx、req 则是 gPRC 方法的前两个参数
info参数包含了当前对应的那个gRPC方法各种信息
handler参数对应当前的gRPC方法函数
*/
func UnaryClientFilter(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Printf("unary filter,method :%v", method)
	return invoker(ctx, method, req, reply, cc, opts...)
}

func StreamClientFilter(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer,
	opts ...grpc.CallOption) (grpc.ClientStream, error) {
	log.Printf("stream filter,method :%v", method)
	return streamer(ctx, desc, cc, method, opts...)
}

// UnaryClientLogging RPC 方法的入参出参的日志输出
func UnaryClientLogging(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Printf("unary gRPC before method: %s req:%v", method, req)
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("unary gRPC after method: %s reply:%v", method, reply)
	return err
}

func StreamClientLogging(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer,
	opts ...grpc.CallOption) (grpc.ClientStream, error) {
	log.Printf("stream gRPC before method: %s", method)
	stream, err := streamer(ctx, desc, cc, method, opts...)
	log.Printf("stream gRPC after method: %s", method)
	return stream, err
}

// UnaryClientRecovery RPC 方法的异常保护和日志输出
func UnaryClientRecovery(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	var err error
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "unary panic err: %v", e)
		}
	}()
	err = invoker(ctx, method, req, reply, cc, opts...)
	return err
}

func StreamClientRecovery(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer,
	opts ...grpc.CallOption) (grpc.ClientStream, error) {
	var err error
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "stream panic err: %v", e)
		}
	}()
	stream, err := streamer(ctx, desc, cc, method, opts...)
	return stream, err
}
