package medium

import (
	"fmt"
	"testing"
)

func Test_listOfDepth(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 8,
				},
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	ret := listOfDepth(root)
	fmt.Println(ret)
}
