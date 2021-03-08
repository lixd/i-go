package lcof

// https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
/*
快指针先走k步，此时快慢指针间距为k
然后快慢指针一起走，等到快指针结束时慢指针距离链表尾部刚好是k
*/
func getKthFromEnd(head *ListNode, k int) *ListNode {

	fast, slow := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
