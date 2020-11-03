package binary_search

// https://leetcode-cn.com/problems/valid-perfect-square
func isPerfectSquare(num int) bool {
	var (
		l, r = 0, num
	)
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
