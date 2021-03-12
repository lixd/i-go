package lcof

// https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
/*
反转链表 将当前节点的next指向前一个节点即可，所以需要存储前一个节点
*/
func reverseList(head *ListNode) *ListNode {
	var (
		prev *ListNode
		curr = head
	)
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
