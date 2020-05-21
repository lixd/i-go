package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"i-go/grpc/stream/proto"
	"i-go/utils"
	"io"
)

const ServerStreamAddress = "localhost:50053"

/*
1. 建立连接 获取client
2. 组装req参数并调用方法获取stream
3. for循环中通过stream.Recv()获取服务端推送的消息
4. err==io.EOF则表示服务端关闭stream了 退出
*/
func main() {
	// 1.建立连接
	conn, err := grpc.Dial(ServerStreamAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 1.获取client
	client := proto.NewServerStreamServerClient(conn)
	// 2.组装req参数
	data := &proto.ServerStreamReq{Data: "1"}
	// 2.调用获取stream
	stream, err := client.ServerStream(context.Background(), data)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "ClientStream Recv error"}).Error(err)
		return
	}

	// 3. for循环获取服务端推送的消息
	for {
		// 3.通过 Recv() 不断获取服务端send()推送的消息
		// 内部也是调用RecvMsg
		data, err := stream.Recv()
		if err != nil {
			// 4. err==io.EOF则表示服务端关闭stream了 退出
			if err == io.EOF {
				fmt.Println("server closed")
				break
			}
			logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "ClientStream Recv error"}).Error(err)
			continue
		}
		fmt.Printf("Recv Data:%v \n", data)
	}
}
