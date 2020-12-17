package main

import (
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "i-go/grpc/stream/proto"
)

type clientStream struct {
	pb.UnimplementedClientStreamServer
}

// ClientStream 客户端流demo
/*
1. for循环中通过stream.Recv()不断接收client传来的数据
2. err == io.EOF表示客户端已经发送完毕关闭连接了,此时在等待服务端处理完并返回消息
3. stream.SendAndClose() 发送消息并关闭连接(虽然客户端流服务器这边并不需要关闭 但是方法还是叫的这个名字)
*/
func (c *clientStream) Sum(stream pb.ClientStream_SumServer) error {
	var sum int64
	// 1.for循环接收客户端发送的消息
	for {
		// 2. 通过 Recv() 不断获取客户端 send()推送的消息
		req, err := stream.Recv() // Recv内部也是调用RecvMsg
		// 3. err == io.EOF表示客户端已经发送完成且关闭stream了
		if err == io.EOF {
			log.Println("client closed")
			// 4.SendAndClose 返回并关闭连接
			// 在客户端发送完毕后服务端即可返回响应
			return stream.SendAndClose(&pb.ClientStreamResp{Sum: sum})
		}
		if err != nil {
			return err
		}
		// 累加求和
		log.Printf("Recved %v", req.Number)
		sum += req.Number
	}
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterClientStreamServer(server, &clientStream{})
	log.Println("Serving gRPC on 0.0.0.0:8081")
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
