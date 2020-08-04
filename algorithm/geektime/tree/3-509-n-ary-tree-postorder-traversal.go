package tree

// https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/
func postorder(root *Node) []int {
	var res []int
	postOrder(root, &res)
	return res
}
func postOrder(root *Node, res *[]int) {
	if root == nil {
		return
	}
	for _, v := range root.Children {
		postOrder(v, res)
	}
	*res = append(*res, root.Val)
}
