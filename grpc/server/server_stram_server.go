package main

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pro "i-go/grpc/proto"
	"i-go/utils"
	"net"
)

const ServerStreamPort = ":50053"

type serverStream struct {
}

var ServerStream = &serverStream{}

// ServerStream
/*
和客户端流相反 是服务端循环发送 然后发送完成后调用
*/
func (server *serverStream) ServerStream(req *pro.ServerStreamReq, stream pro.ServerStreamServer_ServerStreamServer) error {
	fmt.Printf("Recv Client Data %v\n", req.Data)
	for i := 0; i < 5; i++ {
		// 通过 send 方法不断推送数据
		err := stream.Send(&pro.ServerStreamResp{Data: req.Data})
		if err != nil {
			logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "ServerStream send error"}).Error(err)
			return err
		}
	}
	// ? 好像没有close方法 client也能监听到
	return nil
}

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", ServerStreamPort)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer(grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(GenerateInterceptor)))
	pro.RegisterServerStreamServerServer(s, &serverStream{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}

func GenerateInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	logrus.Printf("gRPC method: %s", info.FullMethod)
	err := handler(srv, ss)
	if err != nil {
		logrus.Printf("gRPC err:  %v", err)
	}
	return err
}
