package array_linkedlist_skiplist

// https://leetcode-cn.com/problems/rotate-array/
// 题解 https://leetcode-cn.com/problems/rotate-array/solution/man-hua-san-ci-xuan-zhuan-de-fang-fa-shi-ru-he-x-2/
/*
origin:  1 2 3 4 5 6 7
reverse 7 6 5 4 3 2 1
reverseK 5 6 7 4 3 2 1
reverse n-k 5 6 7 1 2 3 4
*/
func rotate(nums []int, k int) {
	// 先反转整个数组
	reverse(nums)
	// 然后分别反转前K个元素 和后面的元素
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
}
func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}
