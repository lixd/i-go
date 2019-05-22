package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("客户端开始监听...")
	conn, e := net.Dial("tcp", "127.0.0.1:8888")
	if e != nil {
		fmt.Printf("client conn error err=%v \n", e)
		return
	}
	fmt.Println("连接成功", conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		line, e := reader.ReadString('\n')
		if e != nil {
			fmt.Printf("read error err=%v \n", e)
		}
		// 如果用户输入的是 exit 就退出
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("客户端退出")
			break
		}
		_, e = conn.Write([]byte(line + "\n"))
		if e != nil {
			fmt.Printf("conn.Write error err=%v \n", e)
		}

	}

}
