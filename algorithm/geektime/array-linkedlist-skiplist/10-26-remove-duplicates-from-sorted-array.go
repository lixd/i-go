package array_linkedlist_skiplist

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	l, r := 1, 1
	for r < n {
		// 遇到不重复的则移动到 l 指针处，相同的则不管
		if nums[r] != nums[r-1] {
			nums[l] = nums[r]
			l++
		}
		r++
	}
	return l
}
