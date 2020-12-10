package daily

import "math/big"

// uniquePaths https://leetcode-cn.com/problems/unique-paths/
// 到达(i,j)可以从(i-1,j)或(i,j-1)两个地方过来，
// 所以状态转移方程 f(i,j)=f(i−1,j)+f(i,j−1)
// i=0 或 j=0 的时候要忽略 毕竟没有负数的坐标
// 所以将(i,0)和(j,0)都设置为1
func uniquePaths(m, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1 // 处理坐标为0的情况，直接设置为1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1 // 处理坐标为0的情况，直接设置为1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 方法2 组合数学 公式
func uniquePaths2(m, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}
