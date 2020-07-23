package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 需要注意的是，ticker 在不使用时，应该手动 stop，如果不 stop 可能会造成 timer 泄漏。
func main() {
	rand.Seed(time.Now().UnixNano())
	first := rand.Intn(10)
	second := rand.Intn(10)
	fmt.Println("first", first, "second", second)

	//go tickerBad()
	//go tickerGood()
	//go timerBad()
	Load()
	go timerGood()
	select {}
}

type ts struct {
	Host   string
	Status int
}

func Load() {
	res := make(map[string]*ts)
	ts1 := &ts{
		Host:   "127.0.0.1",
		Status: 1,
	}
	res["127.0.0.1"] = ts1
	for _, v := range res {
		fmt.Println(v)
	}
	ts1.Status = 2
	fmt.Println("xxxxxxxxxx")
	for _, v := range res {
		fmt.Println(v)
	}

}

func tickerBad() {
	for {
		select {
		// 错误写法 每次都会 new 一个 ticker
		case t := <-time.Tick(time.Second * 2):
			fmt.Println(t, "time.Tick")
		}
	}
}

func tickerGood() {
	ticker := time.NewTicker(time.Second * 2)
	for {
		select {
		case t := <-ticker.C:
			fmt.Println(t, "time.NewTicker")
		}
	}
}

func timerBad() {
	var ch = make(chan int)
	go func() {
		for {
			ch <- 1
		}
	}()

	for {
		select {
		// 但每次进入 select，time.After 都会分配一个新的 timer。因此会在短时间内创建大量的无用 timer，
		// 虽然没用的 timer 在触发后会消失，但这种写法会造成无意义的 cpu 资源浪费
		case <-time.After(time.Second):
			println("time out, and end")
		case <-ch:
		}
	}
}

func timerGood() {
	var ch = make(chan int)
	go func() {
		for {
			ch <- 1
		}
	}()
	// timer 复用
	timer := time.NewTimer(time.Second)
	for {
		select {
		case <-timer.C:
			timer.Reset(time.Second)
			println("time out, and end")
		case <-ch:
		}
	}
}
