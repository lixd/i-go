package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

// https://blog.wolfogre.com/posts/go-ppof-practice/
func main() {
	var b = Bad{}
	go func() {
		// 模拟一个会一直运行的服务
		for {
			fmt.Println("normal logic")
			b.Range()
		}
	}()
	http.ListenAndServe(":8080", nil)
}

// test
const (
	Ki = 1024
	Mi = Ki * Ki
	Gi = Ki * Mi
)

type Bad struct {
	Buffer [][Mi]byte
}

func (b *Bad) Range() {
	b.Cpu()
	b.Memory()
	b.GC()
	b.Goroutine()
	b.Lock()
	b.Block()
}

//  大量循环 占用 CPU 资源
func (b *Bad) Cpu() {
	fmt.Println("Cpu~~~~~")
	loop := 10000000000
	for i := 0; i < loop; i++ {
		// do nothing
	}
}

// 占用大量内存，一直得不到释放
func (b *Bad) Memory() {
	fmt.Println("Memory~~~~~")
	for len(b.Buffer)*Mi < Gi {
		b.Buffer = append(b.Buffer, [Mi]byte{})
	}
}

// 开辟无用内存并立刻丢弃 频繁触发 gc
func (b *Bad) GC() {
	fmt.Println("GC~~~~~")
	_ = make([]byte, 16*Mi)
}

// 开启大量 goroutine  每个 goroutine 会存在 30s 之后退出（如果子 goroutine 永久阻塞呢？岂不是越堆越多）
func (b *Bad) Goroutine() {
	fmt.Println("Goroutine~~~~~")
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(30 * time.Second)
		}()
	}
}

// 可以看到，这个锁由主协程 Lock，并启动子协程去 Unlock，主协程会阻塞在第二次 Lock 这儿等待子协程完成任务，
// 但由于子协程足足睡眠了一秒，导致主协程等待这个锁释放足足等了一秒钟。
func (b *Bad) Lock() {
	fmt.Println("Lock~~~~~")
	var m sync.Mutex
	m.Lock()
	go func() {
		time.Sleep(time.Second)
		m.Unlock()
	}()
	m.Lock()
}

// 从一个 channel 里读数据时，发生了阻塞，直到这个 channel 在一秒后才有数据读出，这就导致程序阻塞了一秒而非睡眠了一秒。
func (b *Bad) Block() {
	fmt.Println("Block~~~~~")
	<-time.After(time.Second)
}
