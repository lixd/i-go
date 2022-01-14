package bfsdfs

import (
	"fmt"
	"testing"
)

func Test_dfs(t *testing.T) {
	r := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	res := make([][]int, 0)
	res = dfs(r, 0, res)
	fmt.Println(res)
}
