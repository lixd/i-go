package medium

// https://leetcode-cn.com/problems/unique-binary-search-trees/
/*
标签：动态规划
假设 n 个节点存在二叉排序树的个数是 G (n)，令 f(i) 为以 i 为根的二叉搜索树的个数，则
G(n)=f(1)+f(2)+f(3)+f(4)+...+f(n)

当 i 为根节点时，其左子树节点个数为 i-1 个，右子树节点为 n-i，则
f(i)=G(i−1)∗G(n−i) // G(i−1) 为 i-1个节点左子树可以构成的二叉搜索树的个数 G(n−i) 则是右子树个数 乘起来就是以i为根节点的二叉搜索树的个数


综合两个公式可以得到 卡特兰数 公式
G(n)=G(0)∗G(n−1)+G(1)∗(n−2)+...+G(n−1)∗G(0)
*/
func numTrees(n int) int {
	dp := make([]int, n)
	dp[0], dp[1] = 1, 1 // 0和1 边界条件,特殊处理
	for i := 2; i < n; i++ {
		for j := 1; j < i; j++ {
			// f(i)=G(i−1)∗G(n−i)
			dp[i] = dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
