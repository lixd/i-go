package array_linkedlist_skiplist

// https://leetcode-cn.com/problems/linked-list-cycle-ii/

/*
快慢双指针
都从头开始,快指针一次走两步,慢指针一次走一步
如果最后两指针相遇则说明有环
如果相遇之前 快指针已经走到底了说明无环

有环的情况下 快指针从头开始走 一次也走一步 最终肯定会相遇
相遇的位置就是环的起点（为什么呢?参考图）
链长度为 a,环长度为 b。fast 走的路程为 slow 的两倍，假设分别为 f,s 即 f=2s
同时 fast 比 slow多走了 n 个环的长度，即 f = s + nb
以上两式相减得：f = 2nb，s = nb，即fast和slow 指针分别走了 2n，n 个 环的周长。
	然后要走到链表环入口的路程为 a + nb(比如 a+0b 即最短只走 a 步就到了)
当前 slow 已经走了 nb 步了 在走 a 步岂不是就到了入口处？
那么问题来了a 是多少呢？ 所以此时让 slow 原地不动，fast 从 head 开始走，一次走一步 走完链表的长度就到了环的入口。
且 fast slow 又会相遇。
*/
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	var meet *ListNode
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			meet = slow
			break
		}
	}

	// 要用 meet 来判断 有没有 闭环 仅靠 fast.Next != nil 不能保证 slow.Next 一定存在
	if meet == nil {
		return nil
	}

	p := head
	// 快慢指针能相遇说明有环，则头结点和相遇处依次往后移动一个节点，再次相遇处则为环的起始节点
	for p != meet {
		p = p.Next
		meet = meet.Next
	}
	return p
}
