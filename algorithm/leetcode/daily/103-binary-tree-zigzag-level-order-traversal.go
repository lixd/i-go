package daily

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
func zigzagLevelOrder(root *TreeNode) [][]int {
	return bfs(root)
}

func bfs(root *TreeNode) [][]int {
	var (
		result  = make([][]int, 0)
		current = make([]*TreeNode, 0)
	)
	if root == nil {
		return result
	}
	current = append(current, root)
	for level := 0; len(current) > 0; level++ {
		next := make([]*TreeNode, 0)
		ret := make([]int, 0)
		for _, node := range current {
			ret = append(ret, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		// 本质上和层序遍历一样，我们只需要把奇数层的元素翻转即可
		if level%2 == 1 {
			for i, n := 0, len(ret); i < n/2; i++ {
				ret[i], ret[n-1-i] = ret[n-1-i], ret[i]
			}
		}
		current = next
		if len(ret) > 0 {
			result = append(result, ret)
		}
	}
	return result
}
