package binary_search

// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
// 解析 https://leetcode-cn.com/problems/search-in-rotated-sorted-array/solution/jian-dan-luo-lie-mei-yi-chong-qing-kuang-by-liu-zh/
func search(nums []int, target int) int {
	var (
		left, right = 0, len(nums) - 1
	)

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		// 判断是否在前半部分查找
		if (nums[left] <= target && target <= nums[mid]) || (nums[mid] <= nums[right] && (target < nums[mid] || target > nums[right])) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
