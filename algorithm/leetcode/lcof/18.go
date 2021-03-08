package lcof

// https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return head.Next
	}

	h := head // 单独把head记录下来，方便后续直接返回
	for h.Next != nil {
		// 找到待删除节点则删除并退出循环
		if h.Next.Val == val {
			h.Next = h.Next.Next
			break
		}
		// 否则遍历到下一个节点
		h = h.Next
	}
	return head
}
