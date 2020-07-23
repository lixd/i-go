package array_linkedlist_skiplist

// https://leetcode-cn.com/problems/linked-list-cycle/submissions/
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	var (
		slow = head
		fast = head.Next
	)

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
