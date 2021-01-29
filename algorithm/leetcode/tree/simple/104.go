package simple

func maxDepth(root *TreeNode) int {
	return dfs(root)
}

func dfs(root *TreeNode) int {
	var val int
	if root == nil {
		return 0
	}
	// root 节点算一层
	val++
	// 然后加上左子树右子树中最大的层级 就是总层数
	left := dfs(root.Left)
	right := dfs(root.Right)
	val += max(left, right)
	return val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
