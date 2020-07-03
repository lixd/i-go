package simple

/*

110. 平衡二叉树 简单
给定一个二叉树，判断它是否是高度平衡的二叉树。
本题中，一棵高度平衡二叉树定义为：
一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。

*/
// isBalanced
/*
思路：分治法，左边平衡 && 右边平衡 && 左右两边高度 <= 1，
因为需要返回是否平衡及高度，要么返回两个数据，要么合并两个数据，
所以用-1 表示不平衡，>0 表示树高度（二义性：一个变量有两种含义）。
*/
func isBalanced(root *TreeNode) bool {
	/*	if maxDepthDev2(root) == -1 {
			return false
		}
		return true*/
	balanced, _ := maxDepthDev(root)
	return balanced
}

// maxDepthDev
/*
和 max_depth 类似 左右子树分布遍历 3 个条件中有一个不满足则整棵树不平衡 返回 -1
*/
func maxDepthDev(root *TreeNode) (isBalanced bool, dep int) {
	// check
	if root == nil {
		return true, 0
	}
	isLeftBalanced, left := maxDepthDev(root.Left)
	isRightBalanced, right := maxDepthDev(root.Right)

	// 左右子树不平衡 or 高度差大于1 则当前节点也不平衡
	if !isLeftBalanced || !isRightBalanced || left-right > 1 || right-left > 1 {
		return false, 0
	}
	if left > right {
		return true, left + 1
	}
	return true, right + 1
}

// maxDepthDev2 一个变量返回 -1 代表不平衡
func maxDepthDev2(root *TreeNode) int {
	// check
	if root == nil {
		return 0
	}
	left := maxDepthDev2(root.Left)
	right := maxDepthDev2(root.Right)

	// 为什么返回-1呢？（变量具有二义性） [二义性: 一个变量有两种含义]
	// -1 表示不平衡 0 表示无子节点 >0 则说明有多级节点
	/*
		left == -1 左子树不平衡
		right == -1 右子树不平衡
		left-right > 1 || right-left > 1 左右子树高度差 大于 1
	*/
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}
