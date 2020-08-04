package tree

// https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/
// 广度优先
func levelOrder(root *Node) [][]int {
	var res [][]int
	levelOrders(root, &res, 0)
	return res
}

func levelOrders(root *Node, res *[][]int, level int) {
	if root == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, []int{})
	}
	(*res)[level] = append((*res)[level], root.Val)
	for _, v := range root.Children {
		levelOrders(v, res, level+1)
	}
}
