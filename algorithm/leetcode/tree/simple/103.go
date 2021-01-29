package simple

func zigzagLevelOrder(root *TreeNode) [][]int {
	return zigzag(root)
}

func zigzag(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		ans   = make([][]int, 0)
		cur   = make([]*TreeNode, 0)
		level int
	)
	cur = append(cur, root)
	for len(cur) > 0 {
		level++
		next := make([]*TreeNode, 0)
		val := make([]int, 0)
		for _, node := range cur {
			if node == nil {
				continue
			}
			val = append(val, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		cur = next
		if len(val) != 0 {
			// level 记录层数 将偶数层 reverse 即可实现锯齿遍历
			if level%2 == 0 {
				reverse(val)
			}
			ans = append(ans, val)
		}
	}
	return ans
}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}
