package simple

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return isSame(p, q)
}

// BFS
func isSame(p *TreeNode, q *TreeNode) bool {
	// terminator
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	// process
	if p.Val != q.Val {
		return false
	}
	// drill down
	return isSame(p.Left, q.Left) && isSame(p.Right, q.Right)
}
