package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		locker sync.Mutex
		cond   = sync.NewCond(&locker)
	)
	fmt.Printf("%#v \n", cond)
	Print(cond)
	// helloCond()
}
func Print(c *sync.Cond) {
	fmt.Printf("%#v \n", c)
}
func helloCond() {
	var (
		locker sync.Mutex
		cond   = sync.NewCond(&locker)
	)
	for i := 0; i < 40; i++ {
		go func(x int) {
			// wait()方法内部是先释放锁 然后在加锁 所以这里需要先 Lock()
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Wait() // 等待通知,阻塞当前 goroutine
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
	time.Sleep(time.Second)
}
