package difficult

import (
	"fmt"
	"testing"
)

func Test_maxPathSum(t *testing.T) {
	root := &TreeNode{
		Val: 2,
	}
	root.Left = &TreeNode{
		Val: -1,
	}
	//root.Right=&TreeNode{
	//	Val:   2,
	//	Left:  nil,
	//	Right: &TreeNode{
	//		Val:   3,
	//		Left: &TreeNode{
	//			Val:   1,
	//		},
	//	},
	//}
	fmt.Println(maxPathSum(root))
}
