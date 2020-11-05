package binary_search

// https://leetcode-cn.com/problems/sqrtx/
func mySqrt(x int) int {
	var (
		left, right = 0, x
		ans         int
	)
	for left <= right {
		// (right + left) / 2 如果值特别大可能会越界 所以改用下面的方法
		mid := left + (right-left)/2
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
