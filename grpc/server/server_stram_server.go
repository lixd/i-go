package main

import (
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

func (server *serverStream) ServerStream(data *pro.ServerStreamReq, res pro.MyStreamServer_ServerStreamServer) error {

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
	newServer := grpc.NewServer()
	pro.RegisterMyStreamServerServer(newServer, &serverStream{})
	newServer.Serve(lis)
}
