package main

import (
	"context"
	"fmt"
	"time"
)

/*
即时父协程超时返回了，但是子协程仍在继续运行，直到自己退出。
并不能在父协程返回的时候直接 kill 子协程。

goroutine 被设计为不可以从外部无条件地结束掉，只能通过 channel 来与它通信。也就是说，每一个 goroutine 都需要承担自己退出的责任。
(A goroutine cannot be programmatically killed. It can only commit a cooperative suicide.)
相关讨论见:https://github.com/golang/go/issues/32610
摘抄其中几个比较有意思的观点如下：
	1）杀死一个 goroutine 设计上会有很多挑战，当前所拥有的资源如何处理？堆栈如何处理？defer 语句需要执行么？
	2）如果允许 defer 语句执行，那么 defer 语句可能阻塞 goroutine 退出，这种情况下怎么办呢？
因为 goroutine 不能被强制 kill，在超时或其他类似的场景下，为了 goroutine 尽可能正常退出，建议如下：
	1）尽量使用非阻塞 I/O（非阻塞 I/O 常用来实现高性能的网络库），阻塞 I/O 很可能导致 goroutine 在某个调用一直等待，而无法正确结束。
	2）业务逻辑总是考虑退出机制，避免死循环。
	3）任务分段执行，超时后即时退出，避免 goroutine 无用的执行过多，浪费资源。
*/
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// consumer
	go func(ctx context.Context) {
		var ch = make(chan struct{})
		go func() {
			time.Sleep(time.Second * 5)
			ch <- struct{}{}
			// // 套娃？
			// go func() {
			//
			// }()
			// for {
			// 	select {}
			// }
		}()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("child process interrupt...")
				return
			case <-ch:
				return
			}
		}
		// for {
		// 	select {
		// 	case <-ctx.Done():
		// 		fmt.Println("child process interrupt...")
		// 		return
		// 	default:
		// 		fmt.Println("run")
		// 		time.Sleep(time.Second*10)
		// 	}
		// }
	}(ctx)
	defer cancel()
	select {
	case <-ctx.Done():
		time.Sleep(time.Second)
		fmt.Println("main process exit!")
	}
}
