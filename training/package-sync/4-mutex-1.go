package main

import (
	"sync"
	"time"
)

// 都是100ms的周期，但是由于goroutine 1不断的请求锁，可预期它会更频繁的持续到锁。我们基于Go1.8循环了10次，下面是锁的请求占用分布:
// Lock acquired pre goroutine
// g1: 7200216
// g2: 10
// Mutex 被 g1 获取了 700 多万次，而 g2 只获取了 10 次。
// 说明在此案例中，g2 基本上不能获取到锁。
func main() {
	done := make(chan struct{}, 1)
	var mu sync.Mutex
	// goroutine 1
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				mu.Lock()
				time.Sleep(100 * time.Microsecond)
				mu.Unlock()
			}
		}
	}()
	// goroutine 2
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Microsecond)
		mu.Lock()
		mu.Unlock()
	}
	done <- struct{}{}
}
