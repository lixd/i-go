package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	ExampleStdout()
	ExampleIOCopy()
	ExamplePipe()
}

func ExampleIOCopy() (int64, error) {
	return io.Copy(os.Stdout, strings.NewReader("Channels orchestrate mutexes serialize\n"+
		"Cgo is not Go\n"+
		"Errors are values\n"+
		"Don't panic\n"))
}

func ExampleStdout() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize\n",
		"Cgo is not Go\n",
		"Errors are values\n",
		"Don't panic\n",
	}
	for _, p := range proverbs {
		// 因为 os.Stdout 也实现了 io.Writer
		n, err := os.Stdout.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}
}

func ExamplePipe() {
	proverbs := new(bytes.Buffer)
	proverbs.WriteString("Channels orchestrate mutexes serialize\n")
	proverbs.WriteString("Cgo is not Go\n")
	proverbs.WriteString("Errors are values\n")
	proverbs.WriteString("Don't panic\n")

	pipeReader, pipeWriter := io.Pipe()

	// 将 proverbs 写入 pipew 这一端
	go func() {
		defer pipeWriter.Close()
		io.Copy(pipeWriter, proverbs)
	}()

	// 从另一端 piper 中读取数据并拷贝到标准输出
	io.Copy(os.Stdout, pipeReader)
	pipeReader.Close()
}
