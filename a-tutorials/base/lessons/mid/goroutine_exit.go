package main

import (
	"fmt"
	"time"
)

func main() {

}

func forRange(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func forSelectOne(in <-chan int) {
	for {
		select {
		case x, ok := <-in:
			if !ok {
				return
			}
			fmt.Println(x)
		case <-time.After(time.Second * 1):
			fmt.Println("wait...")
		}
	}
}
func forSelectTwo(in <-chan int) {
	for {
		select {
		case x, ok := <-in:
			if !ok {
				// 赋值为 nil 后 select 就不会在当前 case 等待了
				in = nil
			}
			fmt.Println(x)
		case <-time.After(time.Second * 1):
			fmt.Println("wait...")
		}
	}
}

func channel(in <-chan int, stopCh <-chan struct{}) {
	for {
		select {
		case x, ok := <-in:
			if !ok {
				// 赋值为 nil 后 select 就不会在当前 case 等待了
				in = nil
			}
			fmt.Println(x)
		//	同时监听 另一个用于传递 stop 信号的 chan
		case <-stopCh:
			fmt.Println("stop...")
		}
	}
}
