package daily

import (
	"sort"
)

// https://leetcode-cn.com/problems/sliding-window-median/
// 超时
func medianSlidingWindow(nums []int, k int) []float64 {
	var (
		window = make([]int, 0, k)
		ans    = make([]float64, 0, len(nums)-k+1)
	)
	for _, v := range nums {
		window = append(window, v)
		if len(window) > k {
			window = window[1:]
		}
		if len(window) == k {
			f := median(window)
			ans = append(ans, f)
		}
	}
	return ans
}

func median(arr []int) float64 {
	var (
		_len = len(arr)
		a    = make([]int, _len)
	)
	copy(a, arr)
	sort.Ints(a)
	if len(a)%2 == 0 {
		return float64(a[_len/2-1]+a[_len/2]) / 2
	}
	return float64(a[_len/2])
}
