package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"i-go/grpc/stream/proto"
	"i-go/utils"
	"io"
	"net"
	"strings"
)

const ClientStreamPort = ":50054"

type clientStream struct {
}

var ClientStream = &clientStream{}

// ClientStream 客户端流demo
/*
1. for循环中通过stream.Recv()不断接收client传来的数据
2. err == io.EOF表示客户端已经发送完毕关闭连接了,此时服务端需要返回消息
3. stream.SendAndClose() 发送消息并关闭连接(虽然客户端流服务器这边并不需要关闭 但是方法还是叫的这个名字)
*/
func (server *clientStream) ClientStream(stream proto.ClientStreamServer_ClientStreamServer) error {
	list := make([]string, 0)
	// 1.for循环一直接收客户端发送的消息
	for {
		// 2. 通过 Recv() 不断获取服务端send()推送的消息
		data, err := stream.Recv() // Recv内部也是调用RecvMsg
		if err != nil {
			// 3. err == io.EOF表示客户端已经发送完成且关闭stream了
			if err == io.EOF {
				fmt.Println("client closed")
				// 4.SendAndClose 返回并关闭连接
				// 内部调用SendMsg方法 由于这是客户端流 服务端只能发一次 所以没有调用close方法。
				err := stream.SendAndClose(&proto.ClientStreamResp{
					Data: strings.Join(list, ",")})
				if err != nil {
					logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "ClientStream send error"}).Error(err)
					return err
				}
			} else {
				logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "ClientStream Recv error"}).Error(err)
				return err
			}
			return nil
		}
		fmt.Printf("Recv data %v \n", data.Data)
		list = append(list, data.Data)
	}
}

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", ClientStreamPort)
	if err != nil {
		panic(err)
	}
	newServer := grpc.NewServer()
	// 注册server
	proto.RegisterClientStreamServerServer(newServer, &clientStream{})
	_ = newServer.Serve(lis)
}
