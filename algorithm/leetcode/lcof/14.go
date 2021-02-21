package lcof

/*
动态规划
首先定义状态：dp[i]表示长为i的绳子的最大乘积
状态转移方程：dp[i] = max(j段绳子最大值 * i - j段绳子最大值)
比如 dp[4]=dp[2]*dp[2] 或者 dp[4]=dp[1]*dp[3]
*/
func cuttingRope(n int) int {
	if n <= 2 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 1 // 长度为1和2的时候最大值都是1

	// i 表示循环中每次绳子的长度（1、2已经确定了，所以从3开始）
	for i := 3; i < n+1; i++ {
		// j 表示第一次切分后，两段中任意一段的长度，显然另一端的长度就是 i-j
		for j := 1; j < i; j++ {
			// max 函数取得所有切法中的最优解
			// 将长度为i的线段切分后，一段为j，一段为i-j，此时最大值有两种情况
			// 情况一: 认为这已经是最完美的切分方案,不在继续切分了，所以最大值为 j*(i-j)
			// 情况二: 对其中i-j的那一段继续再切分，最大值为 j*dp[i-j]
			// 情况二中为什么不对j这一段进行切分？因为内层循环就是在枚举j的各个取值，所以不需要在切分
			dp[i] = max(dp[i], max(j*dp[i-j], j*(i-j)))
		}
	}

	return dp[n]
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}
