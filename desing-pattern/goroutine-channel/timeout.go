package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	go timeout(ctx, ch)
	time.Sleep(time.Second)
	ch <- struct{}{}
	// do other logic
	time.Sleep(time.Second * 2)
}

func timeout(ctx context.Context, ch <-chan struct{}) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		case v := <-ch:
			fmt.Println(v)
			return
		}
	}
}
