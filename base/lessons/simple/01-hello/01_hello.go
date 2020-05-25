package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	fmt.Println("Hello World~")
}

// path 获取可执行文件路径
func path() string {
	path, _ := exec.LookPath(os.Args[0])
	s, _ := filepath.Abs(path)
	return s
}
