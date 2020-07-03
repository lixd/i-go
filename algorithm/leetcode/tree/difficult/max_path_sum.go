package difficult

/*
124. 二叉树中的最大路径和 困难
给定一个非空二叉树，返回其最大路径和。
本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。

输入: [1,2,3]

       1
      / \
     2   3

输出: 6

思路: 分治法，分为三种情况:
	1) 左子树最大路径和最大
	2) 右子树最大路径和最大
	3) 左右子树最大加根节点最大，
需要保存两个变量：一个保存子树最大路径和，一个保存左右加根节点和，然后比较这个两个变量选择最大值即可
*/
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_, maxPath := cla(root)
	//maxPath := cla2(root)
	return maxPath
}

func cla(root *TreeNode) (singlePath, maxPath int) {
	if root == nil {
		// maxPath 不能返回0 因为节点值有可能是负数 所以只能返回最小值 -(1<<31)
		return 0, -(1 << 31)
	}
	// Divide 左右子树分别计算
	leftSinglePath, leftMaxPath := cla(root.Left)
	rightSinglePath, rightMaxPath := cla(root.Right)
	// 单边最大值
	singlePath = max(leftSinglePath+root.Val, rightSinglePath+root.Val)

	maxPathTmp := max(leftMaxPath, rightMaxPath)
	// 左右子树最大路径 or 左右子树最大加根节点
	maxPath = max(maxPathTmp, leftSinglePath+rightSinglePath+root.Val)
	maxPath = max(singlePath, maxPath)
	//fmt.Println("leftSinglePath:", leftSinglePath, "leftMaxPath:", leftMaxPath)
	//fmt.Println("rightSinglePath:", rightSinglePath, "rightMaxPath:", rightMaxPath)
	//fmt.Println("maxPathTmp", maxPathTmp, "singlePath:", singlePath, "maxPath:", maxPath)

	return singlePath, maxPath
}
func cla2(root *TreeNode) (maxPath int) {
	if root == nil {
		return -(1 << 31)
	}
	// Divide 左右子树分别计算
	leftMaxPath := cla2(root.Left)
	rightMaxPath := cla2(root.Right)

	maxPathTmp := max(leftMaxPath, rightMaxPath)
	// 左右子树最大路径 or 左右子树最大加根节点
	maxPath = max(maxPathTmp, leftMaxPath+rightMaxPath+root.Val)

	return maxPath
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
