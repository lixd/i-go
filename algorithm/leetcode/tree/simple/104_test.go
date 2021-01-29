package simple

import (
	"fmt"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	root := &TreeNode{}
	root.Left = &TreeNode{
		Val: 1,
	}
	root.Right = &TreeNode{
		Val:  2,
		Left: nil,
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(maxDepth(root))
}
