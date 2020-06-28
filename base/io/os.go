package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	file()
}

func file() {
	f, err := os.Create("t.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v \n", f)
	// NewFile 并不是创建文件 而是包装成 File
	// 这里将 标准错误输出 包装成一个 File
	stderr := os.NewFile(uintptr(syscall.Stderr), "/dev/stderr")
	if stderr != nil {
		defer stderr.Close()
		_, _ = stderr.WriteString(
			"The Go language program writes the contents into stderr.\n")
	}
	// 1st: 文件名 2nd: 操作模式 3rd: 权限模式
	os.OpenFile("t.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
}
