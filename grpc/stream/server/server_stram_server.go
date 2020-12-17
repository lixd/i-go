package main

import (
	"log"
	"math"
	"net"

	"google.golang.org/grpc"
	pb "i-go/grpc/stream/proto"
)

type serverStream struct {
	pb.UnimplementedServerStreamServer
}

// Pow ServerStreamDemo 客户端发送一个请求 服务端以流的形式循环发送多个响应
/*
1. 获取客户端请求参数
2. 循环处理并返回多个响应
3. 返回nil表示已经完成响应
*/
func (server *serverStream) Pow(req *pb.ServerStreamReq, stream pb.ServerStream_PowServer) error {
	log.Printf("Recv Client Data %v", req.Number)
	for i := 0; i < 10; i++ {
		// 通过 send 方法不断推送数据
		pow := int64(math.Pow(float64(req.Number), float64(i)))
		err := stream.Send(&pb.ServerStreamResp{Number: pow})
		if err != nil {
			log.Fatalf("Send error:%v", err)
			return err
		}
	}
	// 返回nil表示已经完成响应
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterServerStreamServer(s, &serverStream{})
	log.Println("Serving gRPC on 0.0.0.0:8082")
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
