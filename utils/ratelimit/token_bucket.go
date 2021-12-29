package ratelimit

import (
	"math"
	"sync"
	"time"
)

// 限流算法-令牌桶算法 简单实现

// 定义令牌桶结构
type tokenBucket struct {
	mu       sync.Mutex
	last     time.Time // 上次访问时间
	capacity float64   // 桶的容量（存放令牌的最大量）
	rate     float64   // 令牌放入速度
	tokens   float64   // 当前令牌总量
}

func newTokenBucket(capacity, rate float64) *tokenBucket {
	return &tokenBucket{
		last:     time.Now(),
		capacity: capacity,
		rate:     rate,
	}
}

// allow 判断是否获取令牌
func (t *tokenBucket) allow() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	now := time.Now()
	// 根据时间计算先添加令牌，计算当前时间和上次请求的时间这中间的时候 又增加了多少 token 数。
	t.tokens += now.Sub(t.last).Seconds() * t.rate
	t.tokens = math.Min(t.capacity, t.tokens) // 限制token总数不超过capacity
	t.last = now
	if t.tokens < 1 { // 若桶中一个令牌都没有了，则拒绝
		return false
	}
	// 否则扣除令牌并返回 true
	t.tokens -= 1
	return true
}
