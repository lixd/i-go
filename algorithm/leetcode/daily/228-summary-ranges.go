package daily

import (
	"strconv"
)

// https://leetcode-cn.com/problems/summary-ranges/
func summaryRanges(nums []int) (ans []string) {
	var (
		l = len(nums)
	)
	for i := 0; i < l; {
		left := i
		// 通过这个内层 for 循环使得在一次外层for循环内把相同的数值直接跳过
		for i++; i < l && nums[i-1]+1 == nums[i]; i++ {
		}
		s := strconv.Itoa(nums[left])
		// 然后数值有跳跃则拼接箭头符号
		if left < i-1 {
			s += "->" + strconv.Itoa(nums[i-1])
		}
		ans = append(ans, s)
	}
	return
}

func summaryRanges2(nums []int) []string {
	var (
		ans = make([]string, 0)
	)
	for i := 1; i < len(nums); {
		start := i
		// i++将 for 循环中的 i++ 提前到这里, 防止 i-1 为负数
		i++
		// 通过这个内层 for 循环使得在一次外层for循环内把相同的数值直接跳过
		for i < len(nums) && nums[i]-nums[i-1] == 1 {
			i++
		}
		// 遍历到最后一个数或者数值出现跳跃都要判断一下
		s := strconv.Itoa(nums[start])
		if start < i-1 {
			s += "->" + strconv.Itoa(nums[i-1])
		}
		ans = append(ans, s)
	}
	return ans
}
