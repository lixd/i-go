package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务端开始监听...")
	// "tcp" 表示使用网络协议是 tcp
	// "0.0.0.0:8888" 本地监听8888端口
	listener, err := net.Listen("tcp", "0.0.0.0:7777")
	if err != nil {
		fmt.Printf("listen error err=%v \n", err)
		return
	}
	defer listener.Close()

	for {
		fmt.Println("等待客户端连接")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Accept error err= %v\n", err)
			continue
		}
		fmt.Printf("Accept success clientIP%v \n", conn.RemoteAddr().String())
		go process(conn)
	}

}

func process(conn net.Conn) {
	// 循环接收客户端的连接
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		// 创建一个新的切片
		// 等待客户端通过conn发送信息
		// 如果客户端没有发送，那么协程阻塞在这里
		fmt.Printf("服务器在等待客户端%v发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) // buf
		if err != nil {
			// 出错后直接退出
			fmt.Printf("conn.Read error err= %v\n", err)
			return
		}
		// 显示客户端发送的内容 只显示收到的数据 不显示后面的其他默认数据
		fmt.Printf(string(buf[:n]))
	}
}
