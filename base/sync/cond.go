package main

import (
	"fmt"
	"sync"
	"time"
)

var locker sync.Mutex
var cond = sync.NewCond(&locker)

func main() {
	for i := 0; i < 40; i++ {
		go func(x int) {
			// wait()方法内部是先释放锁 然后在加锁 所以这里需要先 Lock()
			cond.L.Lock()         // 获取锁
			defer cond.L.Unlock() // 释放锁
			cond.Wait()           // 等待通知,阻塞当前 goroutine
			fmt.Println(x)
		}(i)
	}
	for i := 0; i < 30; i++ {
		// 每过50毫秒唤醒一个goroutine
		cond.Signal()
		time.Sleep(time.Millisecond * 50)
	}
	// 剩下10个goroutine一起唤醒
	cond.Broadcast()
	fmt.Println("Broadcast...")
}
