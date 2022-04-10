package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

// go run trace.go 生成 trace.out
// go tool trace trace.out 分析 trace.out 文件
func main() {

	// 创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	// main
	fmt.Println("Hello World")
	var m sync.Map

}
