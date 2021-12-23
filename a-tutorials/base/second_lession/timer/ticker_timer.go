package main

import (
	"fmt"
	"time"
)

/*
timer 一般只用于一次性任务。timer 只会触发一次，虽然可以使用 Reset 重置后再次触发，但是也不推荐。
	time.After 产生大量 timer 的问题：https://www.cnblogs.com/luoming1224/p/11174927.html
ticker 适合用于重复任务，ticker 会一直触发，再不用时需要手动 Stop,否则可能会造成 ticker 泄漏。
*/
func main() {
	// go timerBad()
	// go timerNotRec()
	// go timerGood()

	go tickerBad()
	go tickerGood()
	time.Sleep(time.Second * 3)
}

// timerBad time.After 会创建大量无用的 timer，不建议在循环中调用
func timerBad() {
	for {
		select {
		// 每次进入 执行 time.After 都会分配一个新的 timer。因此会在短时间内创建大量的无用 timer，
		// 虽然没用的 timer 在触发后会消失，但这种写法会造成无意义的 cpu 资源浪费
		case <-time.After(time.Second):
			fmt.Println("timerBad: time.After")
		}
	}
}

// timerNotRec 不推荐使用 Reset 来达到 timer 触发多次的目的
func timerNotRec() {
	timer := time.NewTimer(time.Second)
	for {
		select {
		case <-timer.C:
			fmt.Println("timerNotRec: time.C")
			timer.Reset(time.Second) // reset 后又可以再次触发了,不过不推荐，重复任务更推荐使用 ticker
		}
	}
}

// timerGood timer 只会触发一次，所以只适合用于单次任务
func timerGood() {
	select {
	case <-time.After(time.Second):
		fmt.Println("timerGood: time.After")
	}
}

// tickerBad 和 time.After 存在同样的问题，不推荐在循环中使用 time.Tick
func tickerBad() {
	for {
		select {
		// 不推荐  每次都会 new 一个 ticker
		case <-time.Tick(time.Second * 1):
			fmt.Println("tickerBad time.Tick")
		}
	}
}

// tickerGood 会触发多次适合重复任务，使用 time.NewTicker 以复用 ticker
func tickerGood() {
	ticker := time.NewTicker(time.Second * 1) // 复用 ticker
	defer ticker.Stop()                       // 需要手动停止 ticker，
	for {
		select {
		case <-ticker.C:
			fmt.Println("tickerGood ticker.C")
		}
	}
}
