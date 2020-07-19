package array_linkedlist_skiplist

// https://leetcode.com/problems/swap-nodes-in-pairs
func swapPairs(head *ListNode) *ListNode {
	// 创建一个节点用于记录 head 节点（后续交换中 head 节点会变）
	var dummy = &ListNode{
		Val:  -1,
		Next: head,
	}
	prev := dummy
	for head != nil && head.Next != nil {
		// 找到需要交换的两个节点
		first := head
		second := head.Next
		// 开始交换
		prev.Next = second
		first.Next = second.Next
		second.Next = first
		// 循环到后续两个节点（当前 first 节点已经是第二个节点了）
		prev = first
		head = first.Next
	}
	// 这里直接返回 head 节点
	return dummy.Next
}
