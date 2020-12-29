package daily

// https://leetcode-cn.com/problems/patching-array/
// 题解: https://leetcode-cn.com/problems/patching-array/solution/an-yao-qiu-bu-qi-shu-zu-tan-xin-suan-fa-b4bwr/
func minPatches(nums []int, n int) int {
	var (
		total int // 累加后覆盖的范围
		count int // 需要补充的数字个数
		index int // 访问下标index
	)

	for total < n {
		if index < len(nums) && nums[index] <= total+1 {
			// 如果当前数字在覆盖范围内就不用补充其他数 直接扩大覆盖范围然后继续下一个数
			total += nums[index]
			index++
		} else {
			count++ // 每次都补充最大值 即 total+1 这样覆盖范围直接翻倍
			total = total + (total + 1)
		}
	}
	return count
}
