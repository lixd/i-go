package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("start...")
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	go func() {
		// 	1.
		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("1 done")
				default:
					fmt.Println("1 runing")
					time.Sleep(time.Millisecond * 100)
				}
			}
		}(ctx)
		// 	2.
		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("2 done")
				default:
					fmt.Println("2 runing")
					time.Sleep(time.Millisecond * 100)
				}
			}
		}(ctx)
	}()
	time.Sleep(time.Second * 5)
	cancelFunc()
	time.Sleep(time.Millisecond * 10)
	fmt.Println("end...")
}
