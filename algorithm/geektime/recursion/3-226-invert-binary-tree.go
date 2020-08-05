package recursion

// https://leetcode-cn.com/problems/invert-binary-tree/submissions/
func invertTree(root *TreeNode) *TreeNode {
	invert(root)
	return root
}

func invert(root *TreeNode) {
	// terminator
	if root == nil {
		return
	}
	// process current logic
	root.Right, root.Left = root.Left, root.Right
	// drill down
	if root.Left != nil {
		invert(root.Left)
	}
	if root.Right != nil {
		invert(root.Right)
	}
	// restore current status
}
