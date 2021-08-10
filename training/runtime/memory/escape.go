package main

import (
	"fmt"
	"math/rand"
)

// go build -gcflags -m escape.go

func main() {
	num := getRandom()
	fmt.Println(num)
}

// go:noinline 禁止内联优化，用于测试逃逸分析
func getRandom() *int64 {
	tmp := rand.Int63()
	return &tmp
}
