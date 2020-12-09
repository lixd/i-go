package hotrank

import (
	"math"
)

const (
	// 常数e
	e = 2.71828
	// cd 冷却系数
	cd = 0.192
)

// 笔记链接 https://github.com/lixd/daily-notes/blob/master/DataStructuresandAlgorithms/%E7%83%AD%E5%BA%A6TopN%E6%8E%92%E5%90%8D%E7%AE%97%E6%B3%95.md
// NewtonsLawOfCooling 牛顿冷却定律 公式 T=T0*e^{-α*(t-t0)
/*
latestScore 上次得分
dt 间隔时间（小时）
*/
func NewtonsLawOfCooling(latestScore, dt float64) float64 {
	index := -1 * cd * dt
	score := latestScore * math.Pow(e, index)
	return score
}
