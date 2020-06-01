package linkedlist

func main() {

}

//Definition for singly-linked list.
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
/*
快慢双指针。
快指针先动，领先慢指针k个长度。
快指针到头的时候慢指针刚好在倒数K个
*/
func kthToLast(head *ListNode, k int) int {
	fast := head
	slow := head
	for k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow.Val
}
