package lcof

// https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/
/*
前提:
前序遍历性质： 节点按照 [ 根节点 | 左子树 | 右子树 ] 排序。
中序遍历性质： 节点按照 [ 左子树 | 根节点 | 右子树 ] 排序。
思路:
1)前序遍历的首元素为树的根节点 node 的值。
2)在中序遍历中搜索根节点 node 的索引 ，可将 中序遍历 划分为 [ 左子树 | 根节点 | 右子树 ] 。
3)根据中序遍历中的左 / 右子树的节点数量，可将 前序遍历 划分为 [ 根节点 | 左子树 | 右子树 ] 。
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 根据思路1知道前序遍历第一个元素是根节点
	root := preorder[0]
	for k := range inorder {
		// 这里相等说明inorder中K左边的是左子树 右边的是右子树
		if inorder[k] == root {
			return &TreeNode{
				Val: root, // 根节点是第一个元素可以直接确定
				// 左右子树则递归构建，根据思路2,3和K的值切分出对应的左右子树作为参数传入即可
				Left:  buildTree(preorder[1:k+1], inorder[0:k]),
				Right: buildTree(preorder[k+1:], inorder[k+1:]),
			}
		}
	}
	return nil
}
