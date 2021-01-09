package daily

// https://leetcode-cn.com/problems/partition-list/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// partition 根据值将节点加到不同链表上，最后把两个链表连起来即可。
func partition(head *ListNode, x int) *ListNode {
	var (
		small, large         = &ListNode{}, &ListNode{}
		smallHead, largeHead = small, large
	)

	for head != nil {
		if head.Val >= x {
			large.Next = head
			large = large.Next
		} else {
			small.Next = head
			small = small.Next
		}
		head = head.Next
	}

	// 将两个链表连起来
	large.Next = nil
	small.Next = largeHead.Next
	return smallHead.Next
}
