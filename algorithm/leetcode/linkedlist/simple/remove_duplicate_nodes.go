package main

func main() {

}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
移除重复节点，也就是说一个节点只需要连接一次。
因此只需要用某种方法来判断这个数有没有出现过，如果出现过就不添加。这种时候感觉用哈希是最简单的。
*/
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{}
	ans := newHead
	m := make(map[int]int)

	for head != nil {
		m[head.Val]++
		if m[head.Val] == 1 {
			c := &ListNode{Val: head.Val, Next: nil}
			newHead.Next = c
			newHead = newHead.Next
		}
		head = head.Next
	}
	return ans.Next
}
