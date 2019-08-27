package main

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	featureChan := make(chan float64, 1)
	locusChan := make(chan float64, 1)
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context, cancel context.CancelFunc) {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 200)
			fmt.Println("goroutine1 running", i)
			// 	计算风格模式分数
		}
		if rand.Float64() < 0.8 {
			cancel()
			fmt.Println("goroutine1 cancel")
			featureChan <- 0.0
			return
		} else {
			featureChan <- 1.0
			return
		}
	}(ctx, cancel)
	go func(ctx context.Context, cancel context.CancelFunc) {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 500)
			fmt.Println("goroutine2 running", i)
			// 	计算轨迹相似度分数
		}
		select {
		case <-ctx.Done():
			fmt.Println("goroutine2 cancel")
			locusChan <- 0.0
			return
		}
	}(ctx, cancel)
	a := <-featureChan
	b := <-locusChan
	fmt.Printf("featureChan=%v locusChan=%v \n", a, b)
}

