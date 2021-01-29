package daily

import (
	"sort"
)

// https://leetcode-cn.com/problems/maximum-product-of-three-numbers/
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return max(nums[0]*nums[1]*nums[n-1], nums[n-3]*nums[n-2]*nums[n-1])
}
