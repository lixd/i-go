package main

import (
	"context"
	"fmt"
	"time"
)

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
