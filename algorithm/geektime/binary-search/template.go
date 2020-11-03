package binary_search

// binarySearch 二分查找模板代码（假设array是升序排列的）
func binarySearch(array []int, target int) int {
	var (
		// 1.首先确定左右边界
		left, right = 0, len(array) - 1
	)
	// 一直循环到左下边界大于右边界 还没找到则说明没有这个值
	for left <= right {
		// 取中间值
		mid := (left + right) / 2
		// 相等则直接返回
		if array[mid] == target {
			return mid
		} else if array[mid] < target {
			// 小于则说明目标值在右边 将左边界向右移
			left = right + 1
		} else {
			// 否则说明目标值在左边 将右边界向左移
			right = mid - 1
		}
	}
	return -1
}
