package array_linkedlist_skiplist

// https://leetcode-cn.com/problems/container-with-most-water/
// 暴力计算 // 时间复杂度 O(n^2)
func maxArea(height []int) int {
	var (
		max  int
		area int
	)
	for i := 0; i < len(height); i++ {
		for j := i; j < len(height); j++ {
			area = min(height[i], height[j]) * (j - i)
			if area > max {
				max = area
			}
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 双指针优化 左指针从 0 开始,右指针从 len-1 开始,都往中间靠 每次移动 1
// 具体移动哪个指针则根据 两个指针当前位置的柱子高度来判定 不管移动哪个指针 面积的宽度都会减少1,
// 为了获取最大面积,就只能在高度上有所突破，所以移动当前位置的柱子高度低的那一个指针。
// 当两指针相遇时则停止循环
// 时间复杂度 O(n)
func maxArea2(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	var (
		left      = 0
		right     = len(height) - 1
		area, max int
	)
	for left < right {
		area = min(height[left], height[right]) * (right - left)
		if area > max {
			max = area
		}
		// 当左边指针的值比较小时 就移动左边指针
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

func maxArea3(height []int) int {
	if len(height) < 2 {
		return 0
	}

	var (
		l    = 0
		r    = len(height) - 1
		max  int
		area int
	)
	for l < r {
		area = min(height[l], height[r]) * (r - l)
		if area > max {
			max = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}
