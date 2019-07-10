package main_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func A(ctx context.Context) int {
	ctx = context.WithValue(ctx, "AFunction", "Great")

	go B(ctx)

	select {
	// 监测自己上层的ctx ...
	case <-ctx.Done():
		fmt.Println("A Done")
		return -1
	}
	return 1
}

func B(ctx context.Context) int {
	fmt.Println("A value in B:", ctx.Value("AFunction"))
	ctx = context.WithValue(ctx, "BFunction", 999)

	go C(ctx)

	select {
	// 监测自己上层的ctx ...
	case <-ctx.Done():
		fmt.Println("B Done")
		return -2
	}
	return 2
}

func C(ctx context.Context) int {
	fmt.Println("B value in C:", ctx.Value("AFunction"))
	fmt.Println("B value in C:", ctx.Value("BFunction"))
	select {
	// 结束时候做点什么 ...
	case <-ctx.Done():
		fmt.Println("C Done")
		return -3
	}
	return 3
}

func TestContext(t *testing.T) {
	// 自动取消(定时取消)
	{
		timeout := 10 * time.Second
		ctx, _ := context.WithTimeout(context.Background(), timeout)

		fmt.Println("A 执行完成，返回：", A(ctx))
		select {
		case <-ctx.Done():
			fmt.Println("context Done")
			break
		}
	}
	time.Sleep(20 * time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	go DoCancel(ctx)
	go DoNoting(ctx)
	time.Sleep(time.Second * 10)
	cancel()
	time.Sleep(time.Second * 10)
}

func DoCancel(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("DoCancel: 收到cancel 取消")
			return
		default:
			fmt.Println("DoCancel: 未收到cancel 继续执行")
		}
	}
}
func DoNoting(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("DoNoting: 收到cancel 继续执行")
		default:
			fmt.Println("DoNoting: 未收到cancel 继续执行")
		}
	}
}
