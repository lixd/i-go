// Package ratelimit 速率限制
package ratelimit

import (
	"math"
	"time"
)

// 漏桶算法 伪代码
// 定义漏桶结构
type leakyBucket struct {
	timestamp time.Time // 当前注水时间戳 （当前请求时间戳）
	capacity  float64   // 桶的容量（接受缓存的请求总量）
	rate      float64   // 水流出的速度（处理请求速度）
	water     float64   // 当前水量（当前累计请求数）
}

// 判断是否加水（是否处理请求）
func addWater(bucket leakyBucket) bool {
	now := time.Now()
	// 先执行漏水，计算剩余水量
	leftWater := math.Max(0, bucket.water-now.Sub(bucket.timestamp).Seconds()*bucket.rate)
	bucket.timestamp = now
	if leftWater+1 < bucket.water {
		// 尝试加水，此时水桶未满
		bucket.water = leftWater + 1
		return true
	}
	// 水满了，拒绝加水
	return false
}
