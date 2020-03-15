package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pro "i-go/grpc/proto"
	"net"
	"strconv"
	"time"
)

const ServerStreamPort = ":50053"

type serverStream struct {
}

var ServerStream = &serverStream{}

func (server *serverStream) ServerStream(data *pro.ServerStreamReq, res pro.ServerStreamServer_ServerStreamServer) error {

	for i := 0; i < 20; i++ {
		itoa := strconv.Itoa(i)
		// 通过 send 方法不断推送数据
		err := res.Send(&pro.ServerStreamResp{Data: "count:" + itoa + " client data: " + data.Data})
		if err != nil {
			log.Error(err.Error())
			return err
		}
		time.Sleep(time.Second)
	}
	return nil
}

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", ServerStreamPort)
	if err != nil {
		return
	}
	s := grpc.NewServer(grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(GenerateInterceptor)))
	pro.RegisterServerStreamServerServer(s, &serverStream{})
	s.Serve(lis)
}

func GenerateInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("gRPC method: %s", info.FullMethod)
	err := handler(srv, ss)
	if err != nil {
		log.Printf("gRPC err:  %v", err)
	}
	return err
}
