package main

import (
	"fmt"
	"sync"
)

const (
	Count = 100
)

// 3个函数分别打印 cat、dog、fish,要求每个函数都要起一个goroutine,按照cat、dog、fish顺序打印在屏幕上100次
func main() {
	var (
		wg   sync.WaitGroup
		cat  = make(chan struct{}, 1) // 1.因为最后 fish 会往 cat 中发一个消息,但是 cat 已经打印完成了，所以需要一个缓冲否则会出现死锁。
		dog  = make(chan struct{})
		fish = make(chan struct{})
	)

	go LPrint("cat", cat, dog, &wg)
	go LPrint("dog", dog, fish, &wg)
	go LPrint("fish", fish, cat, &wg)

	cat <- struct{}{} // 2. 因为这里手动触发了一个消息，所以最后 fish 那里会多一个消息出来
	// cat 触发 dog，dog 触发 fish，fish 触发 cat，形成了一个闭环，为了启动整个循环，手动触发了一次 cat 导致最后 fish 无法触发 cat 了，所以 cat 给了一个缓冲
	wg.Add(3)
	wg.Wait()
}

func LPrint(str string, in <-chan struct{}, out chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out)
	for i := 0; i < Count; i++ {
		<-in
		fmt.Println(str)
		out <- struct{}{}
	}
}
