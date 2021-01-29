package medium

import "i-go/algorithm/leetcode/daily"

func generateTrees(n int) []*daily.TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

func helper(start, end int) []*daily.TreeNode {
	if start > end {
		// 没有也要返回一个nil
		return []*daily.TreeNode{nil}
	}
	ret := make([]*daily.TreeNode, 0)
	// 枚举可行根节点
	for i := start; i <= end; i++ {
		// 获得所有可行的左子树集合
		leftTrees := helper(start, i-1)
		// 获得所有可行的右子树集合
		rightTrees := helper(i+1, end)
		// 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				cur := &daily.TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				}
				ret = append(ret, cur)
			}
		}
	}
	return ret
}
