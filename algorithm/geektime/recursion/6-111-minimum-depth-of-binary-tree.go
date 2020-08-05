package recursion

// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree
//
/*
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。 如果左节点或右节点存在一个都不能算叶子节点
每一层一共 4 种情况
 1.根节点为空 则当前层深度为 0
 2.左右节点都为空 则当前层深度为 1，且是最后一层 直接返回
 3.左右节点任意一个不为空 那么都不能算叶子结点 需要继续递归存在的那个子树
 4.左右节点都存在 则两个子树都需要继续递归 且最后取深度最小的一个
*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	}
	if root.Right != nil && root.Left == nil {
		return minDepth(root.Right) + 1
	}
	// 左右节点都存在则递归选最小值
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
