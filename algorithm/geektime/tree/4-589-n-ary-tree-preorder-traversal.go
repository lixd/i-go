package tree

// https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/
func preorder(root *Node) []int {
	var res []int
	nAryPreOrder(root, &res)
	return res
}
func nAryPreOrder(root *Node, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	for _, v := range root.Children {
		nAryPreOrder(v, res)
	}
}
