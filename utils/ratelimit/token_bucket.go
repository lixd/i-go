package ratelimit

import (
	"math"
	"time"
)

// 令牌桶算法 伪代码

// 定义令牌桶结构
type tokenBucket struct {
	timestamp time.Time // 当前时间戳
	capacity  float64   // 桶的容量（存放令牌的最大量）
	rate      float64   // 令牌放入速度
	tokens    float64   // 当前令牌总量
}

// 判断是否获取令牌（若能获取，则处理请求）
func getToken(bucket tokenBucket) bool {
	now := time.Now()
	// 先添加令牌
	leftTokens := math.Max(bucket.capacity, bucket.tokens+now.Sub(bucket.timestamp).Seconds()*bucket.rate)
	bucket.timestamp = now
	if leftTokens < 1 {
		// 若桶中一个令牌都没有了，则拒绝
		return false
	}
	// 桶中还有令牌，领取令牌
	bucket.tokens -= 1
	return true
}
