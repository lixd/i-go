package main

import (
	"log"
	"sync"
	"time"
)

var done = false

func main() {
	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	// write("writer", cond)
	writeSignal("writer1", cond)
	writeSignal("writer2", cond)
	writeSignal("writer3", cond)

	time.Sleep(time.Second)
}

func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		// fmt.Println(name, "wait")
		// NOTE: 这是因为当 Boardcast 唤醒时，有可能其他 goroutine 先于当前 goroutine 唤醒并抢到锁，导致轮到当前 goroutine 抢到锁的时候，可能条件又不再满足了。
		// 因此，需要在 Wait 返回之后再判断一次是否满足条件，最简单的就是直接将条件检查放在 for 循环中。
		// 因为虽然wait之前调用了Lock 但是Wait方法中会调用 Unlock，这中间可能导致done变量被修改 比如在 Read 之后可以把 done 又切换回 false
		c.Wait()
	}
	log.Println(name, "starts reading")
	done = false
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	log.Println(name, "starts writing")
	c.L.Lock()
	done = true
	c.L.Unlock()
	log.Println(name, "wakes all")
	c.Broadcast()
}

func writeSignal(name string, c *sync.Cond) {
	log.Println(name, "starts writing")
	c.L.Lock()
	done = true
	c.L.Unlock()
	c.Signal()
	log.Println(name, "wakes one")
}
