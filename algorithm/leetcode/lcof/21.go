package lcof

// https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
func exchange(nums []int) []int {
	//  双指针，前后遍历
	left, right := 0, len(nums)-1
	for left < right {
		// 正向遍历nums直到nums[left]为偶数，left>right
		// 找到排在最前面的偶数
		for left < right && nums[left]%2 == 1 {
			left++
		}
		// 逆向遍历nums直到nums[right]]为奇数，left>right
		// 找到排在最后的奇数
		for left < right && nums[right]%2 == 0 {
			right--
		}
		// 交换偶数和奇数
		nums[left], nums[right] = nums[right], nums[left]
	}
	return nums
}
