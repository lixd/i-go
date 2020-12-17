package daily

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
/*
dp[n][0]表示第n天无股票状态的收益
dp[n][1]表示第n天有股票状态的收益
手里无股票分为两种情况:
	1.前一天就没有 dp[n][0]=dp[n-1][0]
	2.前一天有,然后当天卖掉了dp[n][0]=dp[n-1][1]+prices[n]-fee
手里有股票分为两种情况:
	1.前一天就有 dp[n][1]=dp[n-1][1]
	2.前一天没有,今天刚买入的 dp[n][1]=dp[n-1][0]-prices[i]
每种情况收益分别取两种情况中的最大值，然后因为手里没有股票的收益肯定是大于有股票时的收益 所以最后结果为dp[n-1][0]
*/
func maxProfit(prices []int, fee int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = maxInt(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = maxInt(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
