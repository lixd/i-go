package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// New一个waitGroup
	waitGroup := sync.WaitGroup{}
	// add 2 表示有两个需要等待
	waitGroup.Add(2)
	for i := 0; i < 2; i++ {
		go func(i int) {
			fmt.Print(i)
			time.Sleep(time.Second)
			// 执行完成后done一个
			defer waitGroup.Done()
		}(i)
	}
	// 程序会阻塞直到计数器减为零
	waitGroup.Wait()
}
