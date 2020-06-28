package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	readWriter()
}

func readWriter() {
	reader := bufio.NewReaderSize(bytes.NewReader([]byte("hello word")), 20)
	// peek 读取后 并不会修改 已读计数
	peek, err := reader.Peek(5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(peek))
	fmt.Printf("%#v \n", reader)
	t := make([]byte, 10)
	_, err = reader.Read(t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(t))
	fmt.Printf("%#v \n", reader)
}
