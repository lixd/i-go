package tree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}
