package daily

// https://leetcode-cn.com/problems/min-cost-climbing-stairs/
/*
上到第n个台阶可以从n-1走一步上来或者n-2走两步上来
f(n)1=f(n-1)+cost[n-1]
f(n)2=f(n-2)+cost[n-2]
每次取最小值即可
fn=min(f(n)1,f(n)2)
*/
func minCostClimbingStairs(cost []int) int {
	// 用dp[i]来表示上到第i阶的花费 不过下标从0开始的 所以需要length+1
	dp := make([]int, len(cost)+1)
	// 然后开始的时候可以选从第0阶或者第1阶开始 所以dp[0]=dp[1]=0
	// 这里就不用初始化(默认值就是0),直接从2开始
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-2]+cost[i-2], dp[i-1]+cost[i-1])
	}
	return dp[len(dp)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
