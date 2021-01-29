package recursion

import "math"

// https://leetcode-cn.com/problems/validate-binary-search-tree/
func isValidBST(root *TreeNode) bool {
	valid := isValid(root, math.MinInt64, math.MaxInt64)
	return valid
}

// isValid 正确写法 每次遍历时带上下一个节点取值的最大值和最小值
func isValid(root *TreeNode, lower, upper int) bool {
	// terminator
	if root == nil {
		return true
	}
	// process
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	// drill down
	return isValid(root.Left, lower, root.Val) && isValid(root.Right, root.Val, upper)
	// restore current status
}

// isValidWrong 错误写法 直接递归判定每个节点是否满足 大于左节点小于右节点
// 然后这样是不行的 比如下面这个例子 按照此方法就是对的 实际上并不是 BST 因为 6 这个节点是错的 右子树任意节点都要大于 根节点
//   10
//  5  15
// x x 6 20
func isValidWrong(root *TreeNode) bool {
	// terminator
	if root == nil {
		return true
	}
	// process
	if root.Left != nil && root.Left.Val > root.Val {
		return false
	}
	if root.Right != nil && root.Right.Val <= root.Val {
		return false
	}
	// drill down
	return isValidWrong(root.Left) && isValidWrong(root.Right)
	// restore current status
}
