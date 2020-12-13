package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net/http"
	"strings"
)

/*
 */
//func main() {
//	auth := auth.Authentication{
//		User:     "17x",
//		Password: "golang",
//	}
//	// grpc.Dial 创建连接 grpc.WithBlock() 阻塞直到连接成功 WithTransportCredentials() tls
//	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithPerRPCCredentials(&auth))
//	if err != nil {
//		panic(err)
//	}
//	defer conn.Close()
//
//	client := proto.NewUserServiceClient(conn)
//	resp, err := client.Create(context.Background(), &proto.UserReq{Name: "17x", Age: "23"})
//	if err != nil {
//		logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "create user  error"}).Error(err)
//	}
//	fmt.Printf("Create User Resp: %v \n", resp.Message)
//}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintln(w, "hello")
	})
	grpcServer := grpc.NewServer()
	_ = http.ListenAndServe(":8080",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor != 2 {
				mux.ServeHTTP(w, r)
				return
			}
			// 根据 Content-Type 来判定走 grpc 还是 http 服务
			if strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				grpcServer.ServeHTTP(w, r) // gRPC Server
				return
			}

			mux.ServeHTTP(w, r)
			return
		}),
	)
}
