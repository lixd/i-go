package hotrange

import (
	"math"
)

const (
	// 常数e
	e = 2.71828
	// cd 冷却系数
	cd = 0.192
)

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
