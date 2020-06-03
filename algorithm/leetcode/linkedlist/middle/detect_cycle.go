package main

/*
面试题 02.08. 环路检测

给定一个有环链表，实现一个算法返回环路的开头节点。
有环链表的定义：在链表中某个节点的next元素指向在它前面出现过的节点，则表明该链表存在环路。

输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。

输入：head = [1,2], pos = 0
输出：tail connects to node index 0
解释：链表中有一个环，其尾部连接到第一个节点。

*/
func main() {

}

/*
1.快慢指针判断是否有环
2.有环则进入第二次循环 寻找入环的第一个节点
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
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
