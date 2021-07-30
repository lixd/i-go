package glimit

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

// 简单的并发限制器 限制 goroutine 数量

type gLimiter struct {
	capacity int64         // 容量
	running  int64         // 正在运行的数量
	duration time.Duration // 间隔
	wg       *sync.WaitGroup
}

func New(size int64, duration time.Duration) *gLimiter {
	if size <= 0 {
		size = 1
	}
	return &gLimiter{
		capacity: size,
		running:  0,
		duration: duration,
		wg:       &sync.WaitGroup{},
	}
}

// Resize 动态调整 pool 大小
func (l *gLimiter) Resize(size int64) {
	if size <= 0 {
		size = 1
	}
	l.capacity = size
}

// Run 增加 running 值,当超过capacity时会被阻塞,可以通过ctx控制超时
func (l *gLimiter) Run(ctx context.Context) bool {
	timer := time.NewTimer(l.duration)
	for {
		select {
		case <-ctx.Done():
			return false
		case <-timer.C:
			if atomic.LoadInt64(&l.running) < atomic.LoadInt64(&l.capacity) {
				atomic.AddInt64(&l.running, 1)
				l.wg.Add(1)
				return true
			}
			timer.Reset(l.duration)
		}
	}
}

func (l *gLimiter) Done() {
	atomic.AddInt64(&l.running, -1)
	l.wg.Done()
}

func (l *gLimiter) Wait() {
	l.wg.Wait()
}
