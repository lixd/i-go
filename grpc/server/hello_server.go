package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "i-go/grpc/proto"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
)

const (
	port = ":50051"
)

// 定义一个结构体
type helloServer struct{}

// 实现proto文件中定义的方法
func (s *helloServer) SayHello(ctx context.Context, in *pb.HelloReq) (*pb.HelloRep, error) {
	return &pb.HelloRep{Message: "Hello " + in.Name}, nil
}

func main() {
	// RPC
	// // 监听
	// lis, err := net.Listen("tcp", port)
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(LoggingInterceptor, RecoveryInterceptor)))
	// // 注册 server
	// pb.RegisterHelloServer(s, &helloServer{})
	// s.Serve(lis)

	// 	提供http接口
	// http://localhost:50054/http -->xxx: gRPC 提供http接口
	mux := GetHTTPServeMux()
	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(LoggingInterceptor)))
	pb.RegisterHelloServer(s, new(helloServer))

	err := http.ListenAndServe("0.0.0.0:50054", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 根据Header判断是否为grpc请求
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			s.ServeHTTP(w, r)
		} else {
			mux.ServeHTTP(w, r)
		}

		return
	}))
	panic(err)
}

func GetHTTPServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/http", func(w http.ResponseWriter, r *http.Request) {
		// 添加逻辑
		w.Write([]byte("xxx: gRPC 提供http接口"))
	})

	return mux
}

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("gRPC method: %s, %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	log.Printf("gRPC method: %s, %v", info.FullMethod, resp)
	return resp, err
}

func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "Panic err: %v", e)
		}
	}()

	return handler(ctx, req)
}
