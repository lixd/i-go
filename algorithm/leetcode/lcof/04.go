package lcof

// 由于二维数组是有规律的，所以可以利用这个规律在降低时间复杂度。
// 从左下角开始寻找target，类似二分查找。
func findNumberIn2DArray(matrix [][]int, target int) bool {
	var (
		r = len(matrix) - 1 // 行
		c = 0               // 列
	)
	if len(matrix) <= 0 || len(matrix[0]) <= 0 {
		return false
	}

	for r >= 0 && c < len(matrix[0]) {
		if matrix[r][c] == target {
			return true
		}
		if matrix[r][c] > target {
			r--
		} else if matrix[r][c] < target {
			c++
		}
	}
	return false
}
