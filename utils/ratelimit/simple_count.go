package ratelimit

import (
	"fmt"
	"sync"
	"time"
)

// Limiter 限流算法 之简单计数器版本
type Limiter struct {
	interval time.Duration // 设置时间窗口大小
	maxCount int           // 窗口内能支持的最大请求数（阈值）
	mu       sync.Mutex    // 并发控制锁
	reqCount int           // 当前窗口请求数（计数器）
}

// NewRateLimiter 构建简单计数器
func NewRateLimiter(interval time.Duration, maxCnt int) *Limiter {
	l := &Limiter{
		interval: interval,
		maxCount: maxCnt,
	}

	go func() {
		ticker := time.NewTicker(interval) // 当达到窗口时间，将计数器清零
		defer ticker.Stop()
		for {
			<-ticker.C
			l.mu.Lock()
			fmt.Println("Reset Count...")
			l.reqCount = 0
			l.mu.Unlock()
		}
	}()

	return l
}

// AllowN 判断是否通过并记录
func (r *Limiter) AllowN(n int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.reqCount+n <= r.maxCount {
		r.reqCount += n
		return true
	}
	return false
}

// ReqCount 对外提供公开方法获取 reqCount
func (r *Limiter) ReqCount() int {
	return r.reqCount
}
