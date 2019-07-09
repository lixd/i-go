package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pro "i-go/grpc/proto"
	"time"
)

const AllStreamAddress = "localhost:50055"

// channel 等待goroutine退出
func main() {
	conn, err := grpc.Dial(AllStreamAddress, grpc.WithInsecure())
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer conn.Close()

	// 通过conn new 一个 client
	client := pro.NewAllStreamServerClient(conn)
	ok := make(chan bool, 2)

	allStr, _ := client.AllStream(context.Background())
	go func() {
		for {
			data, _ := allStr.Recv()
			fmt.Println(data)
		}
		ok <- true
	}()

	go func() {
		for {
			err := allStr.Send(&pro.AllStreamReq{Data: "All Stream From Client"})
			if err != nil {
				log.Error(err.Error())
			}
			time.Sleep(time.Second)
		}
		ok <- true
	}()

	// 让主线程卡在这里 如果两个 goroutine 都结束了 则结束主线程
	for i := 0; i < 2; i++ {
		<-ok
	}

}
