package geektime

import "math"

// Level1 傻递归
func Level1(n int) int {
	if n < 2 {
		return n
	}
	return Level1(n-1) + Level1(n-2)
}

// Level2 Level1 的基础上存储了中间值 减少计算量
var m = make(map[int]int)

func Level2(n int) int {
	res, ok := m[n]
	if ok {
		return res
	}
	if n < 2 {
		return n
	}
	res = Level2(n-1) + Level2(n-2)
	m[n] = res
	return res
}

// Level3 动态规划
// 状态转移方程  DP[n] = DP[n-1] + DP[n-2]
func Level3(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// Level4 数学家给出了通项公式  直接计算..
func Level4(n int) float64 {
	sqrt5 := math.Sqrt(5)
	a := 1 / sqrt5
	b := (1 + sqrt5) * 0.5
	c := (1 - sqrt5) * 0.5
	res := a * (math.Pow(b, float64(n)) - math.Pow(c, float64(n)))
	// res := a *Pow(b, n)-Pow(c, n)
	// 浮点数存在精度问题 结果需要四舍五入
	res = math.Round(res)
	return res
}

// Level5 通过数学矩阵来避免 Level4 中通项公式的浮点数误差
func Level5() {

}
