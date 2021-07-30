package glimit

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func Test_Example(t *testing.T) {
	po := New(1, time.Millisecond)
	println(runtime.NumGoroutine())
	for i := 0; i < 1000; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		run := po.Run(ctx)
		if !run {
			cancel()
			fmt.Println("超时:", i)
			continue
		}
		go func() {
			time.Sleep(time.Millisecond * 10)
			println(runtime.NumGoroutine())
			po.Done()
		}()
		if i%10 == 0 {
			po.Resize(po.capacity + 2)
		}
		if i%100 == 0 {
			po.Resize(2)
		}
		cancel()
	}
	po.Wait()
	println(runtime.NumGoroutine())
}
