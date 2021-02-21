package lcof

import "math"

// 暴力解法 遍历找出最小值 O(n)
func minArray(numbers []int) int {
	var min = math.MaxInt64
	for _, v := range numbers {
		if v < min {
			min = v
		}
	}
	return min
}

// 二分法 O(logn)
/*
如果是升序数组 那么肯定有 numbers[left] <= numbers[mid] <= numbers[right], 但是旋转后，有时就违反了上述条件，我们分别作分析：
需要注意的是:题目说的是把从最小值开始的那一段旋转到数组尾部，所以实际要找的是左右两个小数组中，右边数组的最小值。

1): numbers[mid] > numbers[right]
中间端点大于右侧端点，说明mid还处于左边数组，因此左端点要右移：left = mid + 1
2): numbers[mid] < numbers[right]
中间端点小于右侧端点，说明mid已经处于右边数组了,因为是升序数组，所以最小值在mid的前面，所以右端点需要左移：right = mid
3): numbers[mid] == numbers[right]
中间端点等于右侧端点，说明right不是最小值或者不是唯一的最小值（可能有重复元素），只能慢慢像中间移动：right = right - 1
*/
func minArray2(numbers []int) int {
	var (
		left  = 0
		right = len(numbers) - 1
	)
	for left < right {
		mid := left + (right-left)>>2
		// 下一轮搜索区间是 [mid + 1, right]
		if numbers[mid] > numbers[right] {
			left = mid + 1
			// mid 的右边一定不是最小数字，mid 则有可能是，下一轮搜索区间是 [left, mid]
		} else if numbers[mid] < numbers[right] {
			right = mid
		} else {
			// 此时 numbers[mid] = numbers[right]
			// 只能把 right 排除掉，下一轮搜索区间是 [left, right - 1]
			right = right - 1
		}
	}
	return numbers[left]
}
