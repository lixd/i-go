package medium

import "math"

func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}

// 不能直接递归判断左右子树是否是BST
func isValid(root *TreeNode, lower, upper int) bool {
	// terminator
	if root == nil {
		return true
	}
	// 判断当前节点是否满足条件
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	//  drill down 继续判断左右子树
	return isValid(root.Left, lower, root.Val) && isValid(root.Right, root.Val, upper)
}
