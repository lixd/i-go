package lcof

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
// 全部存到数组中去，然后反转数组 这样会比递归快一些
func reversePrint(head *ListNode) []int {
	var ans = make([]int, 0)
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}
	reverse(ans)
	return ans
}
func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}
