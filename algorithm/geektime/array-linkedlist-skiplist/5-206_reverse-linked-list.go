package array_linkedlist_skiplist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	var (
		prev *ListNode = nil
		curr           = head
		tmp  *ListNode
	)
	for curr != nil {
		// 保存下一个节点 备用
		tmp = curr.Next
		// 反转当前节点
		curr.Next = prev
		// 移动到下一个位置
		prev = curr
		curr = tmp
	}
	return prev
}
