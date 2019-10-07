package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// NewReader
	r := strings.NewReader("abcdefghijklmn")
	// Len Size
	fmt.Println("Len ", r.Len(), "Size ", r.Size())
	// Read 长度为1时=ReadByte
	b := make([]byte, 1)
	r.Read(b)
	r.ReadByte()
	fmt.Println("Read ", b)
	// ReadByte
	bunRead := make([]byte, 1)
	r.UnreadByte()
	r.Read(bunRead)
	fmt.Println("UnreadByte ", bunRead)
	// 	ReadAt
	bat := make([]byte, 10)
	r.ReadAt(bat, 1)
	fmt.Println("ReadAt ", bat)
	// Seek
	bSeek := make([]byte, 10)
	i, _ := r.Seek(-1, 1)
	fmt.Println("Seek Index ", i)
	r.Read(bSeek)
	fmt.Println("Seek ", bSeek)

	// 	WriteTo
	file, err := os.OpenFile("/WriteTo.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("OpenFile err ", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	r.WriteTo(writer)
	writer.Flush()

	// 	Reset()
	r.Reset("newreader")
	fmt.Println("Len ", r.Len(), "Size ", r.Size())
}
