package greedy

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/description/
func maxProfit(prices []int) int {
	var maxProfit int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			// 所有的收益点累加起来 最后收益一定是最大的
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}
