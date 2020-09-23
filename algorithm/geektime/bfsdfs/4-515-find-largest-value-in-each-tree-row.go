package bfsdfs

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	return bfsFind(root)
}

func bfsFind(root *TreeNode) (result []int) {
	var (
		queue = make([]*TreeNode, 0)
	)
	queue = append(queue, root)
	for len(queue) > 0 {
		next := make([]*TreeNode, 0)
		var max = math.MinInt64
		for _, v := range queue {
			if v == nil {
				continue
			}
			if v.Left != nil {
				next = append(next, v.Left)
			}
			if v.Right != nil {
				next = append(next, v.Right)
			}
			if v.Val > max {
				max = v.Val
			}
		}
		if max != math.MinInt64 {
			result = append(result, max)
		}
		queue = next
	}
	return
}
