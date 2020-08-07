package recursion

import "i-go/algorithm/geektime/tree"

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree
func lowestCommonAncestor(root, p, q *tree.TreeNode) *tree.TreeNode {
	// terminator
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	// 	process
	// drill down
	left := lowestCommonAncestor(root.Left, q, p)
	right := lowestCommonAncestor(root.Right, q, p)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
	// restore status
}
