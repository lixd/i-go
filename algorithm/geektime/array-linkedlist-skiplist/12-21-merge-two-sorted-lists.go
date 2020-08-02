package array_linkedlist_skiplist

// https://leetcode-cn.com/problems/merge-two-sorted-lists
// 迭代的方式 比较两个链表当前节点的大小 小的就添加到新链表且往后移动一位
// 最后可能会剩下一个节点没有添加进去 所以需要单独判定一次
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var hair = &ListNode{
		Val:  0,
		Next: nil,
	}
	result := hair
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			hair.Next = l1
			l1 = l1.Next
		} else {
			hair.Next = l2
			l2 = l2.Next
		}
		hair = hair.Next
	}
	// 添加漏掉的节点
	if l1 != nil {
		hair.Next = l1
	}
	if l2 != nil {
		hair.Next = l2
	}
	return result.Next
}
