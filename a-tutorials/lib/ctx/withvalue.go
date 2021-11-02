package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// 通过 ctx 进行超时控制的同时，在 ctx 中存放 traceId 进行链路追踪。
func main() {
	// f2 需要 10ms 这里超时设置1ms 肯定会触发超时
	withTimeout, cancel := context.WithTimeout(context.Background(), time.Millisecond*1)
	defer cancel()
	ctx := context.WithValue(withTimeout, "traceId", "id12345")
	r := f1(ctx)
	fmt.Println("r:", r)
	time.Sleep(time.Second) // 等待 ctx 超时
}

func f1(ctx context.Context) int {
	fmt.Println("f1 traceId:", fromCtx(ctx))
	var ret = make(chan int, 1)
	go f2(ctx, ret)
	r1 := rand.Intn(10)
	fmt.Println("r1:", r1)
	select {
	case <-ctx.Done():
		fmt.Println("f1 ctx timeout")
		return r1
	case r2 := <-ret:
		return r1 + r2
	}
}

func f2(ctx context.Context, ret chan int) {
	fmt.Println("f2 traceId:", fromCtx(ctx))
	// sleep 模拟耗时逻辑
	time.Sleep(time.Millisecond * 10)
	r2 := rand.Intn(10)
	fmt.Println("r2:", r2)
	ret <- r2
}

func fromCtx(ctx context.Context) string {
	return ctx.Value("traceId").(string)
}
