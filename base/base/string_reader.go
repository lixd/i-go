package main

import (
	"bufio"
	"fmt"
	"io"
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
	// 	Example
	ExampleRead()
}

func ExampleRead() {
	reader := strings.NewReader("Clear is better than clever")
	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF:", n)
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(n, string(p[:n]))
	}
}

// -------------自定义Reader Begin----------
type alphaReader struct {
	// 资源
	src string
	// 当前读取到的位置
	cur int
}

// 创建一个实例
func newAlphaReader(src string) *alphaReader {
	return &alphaReader{src: src}
}

// 过滤函数
func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

// 实现Read 方法
func (a *alphaReader) Read(p []byte) (int, error) {
	// 当前位置 >= 字符串长度 说明已经读取到结尾 返回 EOF
	if a.cur >= len(a.src) {
		return 0, io.EOF
	}

	// lastLen 是剩余未读取的长度
	lastLen := len(a.src) - a.cur
	// canReadLen 本次能够读取的长度
	var canReadLen int
	if lastLen >= len(p) {
		// 剩余长度超过缓冲区大小，说明本次可完全填满缓冲区
		canReadLen = len(p)
	} else if lastLen < len(p) {
		// 剩余长度小于缓冲区大小，使用剩余长度输出，缓冲区不补满
		canReadLen = lastLen
	}

	buf := make([]byte, canReadLen)
	// 当前buf的索引
	var curIndexBuf int
	for curIndexBuf < canReadLen {
		// 每次读取一个字节，执行过滤函数
		if char := alpha(a.src[a.cur]); char != 0 {
			buf[curIndexBuf] = char
		}
		curIndexBuf++
		a.cur++
	}
	// 将处理后得到的 buf 内容复制到 p 中
	copy(p, buf)
	return curIndexBuf, nil
}

// -------------自定义Reader End----------

// --------------Reader组合 Begin---------------
type alphaReaderCom struct {
	// alphaReader 里组合了标准库的 io.Reader
	reader io.Reader
}

func newAlphaReaderCom(reader io.Reader) *alphaReaderCom {
	return &alphaReaderCom{reader: reader}
}

func (a *alphaReaderCom) ReadCom(p []byte) (int, error) {
	// 这行代码调用的就是 io.Reader
	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}
	// 然后加入自定义逻辑
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := alpha(p[i]); char != 0 {
			buf[i] = char
		}
	}

	copy(p, buf)
	return n, nil
}

// --------------Reader组合 End---------------

// -----------------自定义Writer Begin-----------------
type chanWriter struct {
	// ch 实际上就是目标资源
	ch chan byte
}

func newChanWriter() *chanWriter {
	return &chanWriter{make(chan byte, 1024)}
}

func (w *chanWriter) Chan() <-chan byte {
	return w.ch
}

func (w *chanWriter) Write(p []byte) (int, error) {
	n := 0
	// 遍历输入数据，按字节写入目标资源
	for _, b := range p {
		w.ch <- b
		n++
	}
	return n, nil
}

func (w *chanWriter) Close() error {
	close(w.ch)
	return nil
}

// -----------------自定义Writer End-----------------
