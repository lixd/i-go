package main

import (
	"log"
	"math/rand"
	"net"
	"strconv"
)

func main() {
	//建立socket，监听端口  第一步:绑定端口
	netListen, err := net.Listen("tcp", "localhost:9800")
	if err != nil {
		panic(err)
	}
	//defer延迟关闭改资源，以免引起内存泄漏
	defer netListen.Close()
	Log("Waiting for clients")
	for {
		//第二步:等待连接
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		Log(conn.RemoteAddr().String(), " tcp connect success")
		// 使用goroutine来处理用户的请求
		go handleConnection(conn)
	}
}

//处理连接
func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048)
	for {
		//第三步:读取从该端口传来的内容
		n, err := conn.Read(buffer)
		// 随便写点数据回去
		words := "golang socket server : " + strconv.Itoa(rand.Intn(100))
		conn.Write([]byte(words))
		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		Log(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
	}
}

//log输出
func Log(v ...interface{}) {
	log.Println(v...)
}
