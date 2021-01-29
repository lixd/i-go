package simple

func isSymmetric(root *TreeNode) bool {
	return symmetric(root, root)
}

func symmetric(x, y *TreeNode) bool {
	// 如果两个子树都为空指针，则它们相等或对称
	if x == nil && y == nil {
		return true
	}
	// 如果两个子树只有一个为空指针，则它们不相等或不对称
	if x == nil || y == nil {
		return false
	}
	// 如果两个子树根节点的值不相等，则它们不相等或不对称
	if x.Val != y.Val {
		return false
	}
	return symmetric(x.Left, y.Right) && symmetric(x.Right, y.Left)
}
