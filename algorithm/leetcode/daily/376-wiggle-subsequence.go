package daily

// https://leetcode-cn.com/problems/wiggle-subsequence/

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	up, down := 1, 1
	for i := 1; i < n; i++ {
		// 当前元素小于前一个元素说明是下降序列否则为上升序列 相等则直接排除掉（不计入长度）
		if nums[i] > nums[i-1] {
			up = max(up, down+1)
		} else if nums[i] < nums[i-1] {
			down = max(up+1, down)
		}
	}
	return max(up, down)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
