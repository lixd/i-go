package simple

func levelOrder(root *TreeNode) [][]int {
	return bfs(root)
}
func bfs(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		cur = make([]*TreeNode, 0)
		ans = make([][]int, 0)
	)
	cur = append(cur, root)
	for len(cur) > 0 {
		// next 存储下一层的节点
		next := make([]*TreeNode, 0)
		val := make([]int, 0, 4)
		for _, v := range cur {
			val = append(val, v.Val)
			if v.Left != nil {
				next = append(next, v.Left)
			}
			if v.Right != nil {
				next = append(next, v.Right)
			}
		}
		if len(val) > 0 {
			ans = append(ans, val)
		}
		cur = next
	}
	return ans
}
