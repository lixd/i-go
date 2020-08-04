package tree

// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
func preorderTraversal(root *TreeNode) []int {
	var res []int
	preOrder(root, &res)
	return res
}
func preOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preOrder(root.Left, res)
	preOrder(root.Right, res)
}
