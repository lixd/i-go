package greedy

import "sort"

// https://leetcode-cn.com/problems/coin-change/
// dp 贪心算法好像并不是很科学 给出的硬币如果不是倍数关系则最后得不到最优解 如【10,7,1】构成14 贪心算法结果一般是10然后加多个1 实际最优则是两个7
/*
假设最小组合中最后一枚硬币是c：

c 可能是 coins 中任何一个；
去除 c 后剩下的部分，一定也是当前总额的一个最小组合（否则加上c不可能构成最小组合）
或者用以下思路：
如果 dp[i] 表示组成金额i的最小组合，dp[i]+1 一定是组成金额 i+c 的最小组合。
 动态规划的方程
dp[i]=min(dp[i-x],dp[i-y],dp[i-z])+1  x y z 为金额面值
*/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if i >= coin {
				if dp[i-coin]+1 < dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}
	if dp[amount] < amount+1 {
		return dp[amount]
	}
	return -1
}

// 贪心算法
type solution struct {
	max    int
	amount int
	coins  []int
}

func coinChange2(coins []int, amount int) int {
	s := newSolution(coins, amount)
	s.DFS(0, len(coins)-1, 0)
	return s.max
}
func newSolution(coins []int, amount int) *solution {
	sort.Ints(coins)
	return &solution{-1, amount, coins}
}
func (s *solution) DFS(currVal, index, currCount int) {
	if s.amount == currVal {
		if currCount < s.max || s.max < 0 {
			s.max = currCount
			return
		}
	} else if currVal > s.amount {
		return
	}
	for i := index; i >= 0; i-- {
		if currVal+s.coins[i] > s.amount {
			continue
		}
		if s.max > 0 && (s.amount-currVal)/s.coins[i] > s.max-currCount {
			break
		}
		currVal += s.coins[i]
		s.DFS(currVal, i, currCount+1)
		currVal -= s.coins[i]
	}
}
