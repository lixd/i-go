// Package ratelimit 速率限制
package ratelimit

import (
	"math"
	"time"
)

// 限流算法-漏桶算法 简单实现

type leakyBucket struct {
	last     time.Time // 当前注水时间戳 （当前请求时间戳）
	capacity float64   // 桶的容量（接受缓存的请求总量）
	rate     float64   // 水流出的速度（处理请求速度）
	water    float64   // 当前水量（当前累计请求数）
}

func newLeakyBucket(capacity, rate float64) *leakyBucket {
	return &leakyBucket{
		last:     time.Now(),
		capacity: capacity,
		rate:     rate,
	}
}

// allow 判断能否通过
func (l *leakyBucket) allow() bool {
	now := time.Now()
	// 先执行漏水，同样是根据时间计算剩余水量
	l.water -= now.Sub(l.last).Seconds() * l.rate
	l.water = math.Max(0, l.water) // 限制 水被扣成负值
	l.last = now
	if l.water+1 <= l.capacity { // 若漏桶没有满,则加水，
		l.water += 1
		return true
	}
	// 水满了，拒绝加水
	return false
}
