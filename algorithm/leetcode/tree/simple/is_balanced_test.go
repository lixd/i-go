package simple

import (
	"fmt"
	"testing"
)

func Test_isBalanced(t *testing.T) {
	root := &TreeNode{}
	root.Left = &TreeNode{
		Val: 1,
	}
	root.Right = &TreeNode{
		Val:  2,
		Left: nil,
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
		},
	}
	fmt.Println(isBalanced(root))
}
