package recursion

// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
func maxDepth(root *TreeNode) int {
	level := depth(root, 0)
	return level
}

func depth(root *TreeNode, level int) int {
	// 	terminator
	if root == nil {
		return level
	}
	// process
	// restore current status
	return max(depth(root.Left, level+1), depth(root.Right, level+1))
}

// depth2 与 depth 相比只是把每层的 level+1 放在的 return那个地方
func depth2(root *TreeNode) int {
	// 	terminator
	if root == nil {
		return 0
	}
	// process
	// restore current status
	return max(depth2(root.Left), depth2(root.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
