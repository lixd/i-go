package main

import "fmt"

/*
这两个 API 有什么区别？
// ListDirectory return the contents of dir
func ListDirectory(dir string) ([]string, error)

// ListDirectory returns a channel over which
// directory entries will be published. When the list
// of entries is exhausted, the channel will be closed.
func ListDirectory(dir string) chan string

// 推荐做法，使用回调的方式，由调用者来控制并发逻辑 标准库 `filepath. WalkDir`也是类似的模型
func ListDirectory(dir string,fn func(string)) chan string
*/

// leak is a buggy function. It launches a goroutine that
// blocks receiving from a channel. Nothing will ever be
// sent on that channel and the channel is never closed SO
// that goroutine will be blocked forever,
func leak() {
	ch := make(chan int)
	go func() {
		val := <-ch
		fmt.Println("We received a value:", val)
	}()
}
