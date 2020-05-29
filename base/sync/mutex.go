package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	defer func() {
		m.Unlock()
		wg.Done()
	}()
	// 由于加锁了 所以同时只会有一个goroutine在执行这段代码
	// 不会出现值丢失的问题
	x = x + 1
}

func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		// 1000个goroutine并发修改x的值
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
