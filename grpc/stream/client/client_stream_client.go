package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"i-go/grpc/stream/proto"
	"i-go/utils"
	"strconv"
)

const ClientStreamAddress = "localhost:50054"

/*
1. 建立连接并获取client
2. 通过stream.Send()循环发送消息
3. 发送完成后通过stream.CloseAndRecv() 关闭steam并接收服务端返回结果
*/
func main() {
	// 1.建立连接
	conn, err := grpc.Dial(ClientStreamAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 1. 获取client
	client := proto.NewClientStreamServerClient(conn)

	// 获取stream
	stream, err := client.ClientStream(context.Background())
	if err != nil {
		logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "ClientStream Get Client error"}).Error(err)
		return
	}

	for i := 0; i < 5; i++ {
		// 2.通过 send 方法不断推送数据到server
		// send方法内部调用的就是SendMsg方法
		err := stream.Send(&proto.ClientStreamReq{Data: strconv.Itoa(i)})
		if err != nil {
			logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "ClientStream send error"}).Error(err)
			return
		}
	}
	// 3. CloseAndRecv关闭连接并接收服务端返回结果(服务端则根据err==io.EOF来判断client是否关闭stream)
	// CloseAndRecv内部也只是调用了 CloseSend方法和RecvMsg方法
	resp, err := stream.CloseAndRecv()
	if err != nil {
		logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "ClientStream Recv error"}).Error(err)
		return
	}
	fmt.Printf("Recv %v \n", resp.Data)
}
