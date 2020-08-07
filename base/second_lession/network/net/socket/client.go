package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9800")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	var data string
	_, err = fmt.Scan(&data)
	_, err = conn.Write([]byte(data))
	if err != nil {

	}
	buffer := make([]byte, 2048)
	var n int
	for {
		n, _ := conn.Read(buffer)
		if n != 0 {
			break
		}
	}
	fmt.Println(conn.RemoteAddr().String(), "read data string:\n", string(buffer[:n]))
}
