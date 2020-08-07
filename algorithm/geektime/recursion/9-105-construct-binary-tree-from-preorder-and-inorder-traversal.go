package recursion

import "i-go/algorithm/geektime/tree"

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
func buildTree(preorder []int, inorder []int) *tree.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 前序遍历顺序为根-左-右 所以第一个元素为 root
	root := &tree.TreeNode{preorder[0], nil, nil}
	// 然后这里找到中序遍历中 根节点所在的位置
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	// 由这个位置对中序遍历数组进行切分 i之前的为左子树 i之后的为右子树0
	// 同理前序遍历中 0为根 那么 0到i+1 之间的就是左子树 i+1之后的为右子树
	// 然后分别对左右子树进行递归
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
