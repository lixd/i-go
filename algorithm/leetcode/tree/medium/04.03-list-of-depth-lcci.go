package medium

// https://leetcode-cn.com/problems/list-of-depth-lcci/

// listOfDepth bfs 广度优先搜索 然后把每层的节点构造成链表即可
func listOfDepth(tree *TreeNode) []*ListNode {
	return bfs(tree)
}

func bfs(root *TreeNode) []*ListNode {
	var (
		current = make([]*TreeNode, 0) // 当前层
		next    = make([]*TreeNode, 0) // 下一层
		ret     = make([]*ListNode, 0)
	)
	current = append(current, root)
	for len(current) > 0 {
		node := new(ListNode)
		head := node
		for i := 0; i < len(current); i++ {
			v := current[i]
			if i == 0 {
				node = &ListNode{Val: v.Val}
			} else {
				node.Next = &ListNode{Val: v.Val}
				node = node.Next
			}
			if v.Left != nil {
				next = append(next, v.Left)
			}
			if v.Right != nil {
				next = append(next, v.Right)
			}
		}
		ret = append(ret, head)
		current = next
		next = nil
	}
	return ret
}
