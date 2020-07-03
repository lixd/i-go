package simple

func main() {

}

func maxDepth(root *TreeNode) int {
	// 为空则返回0
	/*	每次都只会返回 0,如果不为 nil 则继续迭代 直到为 nil
		最后从叶子节点一直往回走 每层 +1 最后就能得到*/
	if root == nil {
		// fmt.Println("nil ")
		return 0
	}

	// divide：分左右子树分别计算
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	// conquer：合并左右子树结果
	if left > right {
		// +1 是因为还有 root 节点这一层
		return left + 1
	}

	return right + 1
}
