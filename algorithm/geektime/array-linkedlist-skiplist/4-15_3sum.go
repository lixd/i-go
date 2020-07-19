package array_linkedlist_skiplist

import (
	"sort"
)

// https://leetcode-cn.com/problems/3sum/
// 1.暴力求解：三重循环 O(n^3)
// 2.hash表
// 3.左右推进
/*
首先进行一次遍历 选当前元素为 k
然后指定左右指针(l,r) 分别为 k+1 和 len(nums)-1
然后不断移动左右指针 尝试 k+l+r 是否 =0
如果小于0 则 l++(因为数组有序的,左指针右移后增大 k+l+r 可能就等于0了)
如果大于0 则 r-- 右指针左移 整体会减少 k+l+r 也有能等于0
当 l>=r 后结束本轮循环 k+1 左右指针重新赋值为k+1 和 len(nums)-1
一直到结束
最后则需要过滤掉重复的
首先是要过滤掉 k 相同的 `if i == 0 || nums[i] > nums[i-1]`
然后过滤掉 左指针与上一个值相同的情况 右指针与上一个值相同的情况
`for l < r && nums[l] == nums[l-1] {
	l++
}
// 跳过相同的
for r > l && nums[r] == nums[r+1] {
	r--
}`
*/
func threeSum(nums []int) [][]int {
	var ret [][]int
	// 首先要排下序，就变成有序数组r、l指针扫描问题了
	sort.Ints(nums)
	var l, r int
	for i := 0; i < len(nums)-2; i++ {
		// nums [i]> nums[i-1]过滤掉重复搜索的情况
		if i == 0 || nums[i] > nums[i-1] {

			l, r = i+1, len(nums)-1
			for l < r {
				if nums[l]+nums[r]+nums[i] == 0 {
					// 找到一个匹配，记录并调整搜索的左右边界
					ret = append(ret, []int{nums[i], nums[l], nums[r]})
					l++
					r--
					// 跳过相同的
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					// 跳过相同的
					for r > l && nums[r] == nums[r+1] {
						r--
					}
				} else {
					if nums[l]+nums[r]+nums[i] > 0 {
						r--
					} else {
						l++
					}
				}
			}
		}
	}
	return ret
}

func threeSum2(nums []int) [][]int {
	var l, r int
	var ret [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		// nums[i] > nums[i-1] 过滤掉重复的
		if i == 0 || nums[i] > nums[i-1] {
			l, r = i+1, len(nums)-1
			for l < r {
				if nums[l]+nums[r]+nums[i] == 0 {
					// 找到一个匹配，记录并调整搜索的左右边界
					ret = append(ret, []int{nums[l], nums[r], nums[i]})
					l++
					r--
					// 过滤掉重复的 用 for 循环以便过滤掉多个相同的值
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					// 过滤掉重复的
					for r > l && nums[r] == nums[r+1] {
						r--
					}
				} else {
					if nums[l]+nums[r]+nums[i] > 0 {
						r--
					} else {
						l++
					}
				}
			}
		}
	}
	return ret
}

func threeSum3(nums []int) [][]int {
	var (
		l   int
		r   int
		res [][]int
	)
	sort.Ints(nums)
	// 不能重复 所以到  len(nums)-2 即可
	for i := 0; i < len(nums)-2; i++ {
		// i 可以直接过滤掉重复的情况 （i=0 特殊情况 不判断）
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r = i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r]+nums[i] == 0 {
				res = append(res, []int{nums[l], nums[r], nums[i]})
				l++
				r--
				// l r 要先找到匹配的组合之后再过滤重复的。。
				for l > 1 && l < r && nums[l] == nums[l-1] {
					l++
				}
				for r < len(nums)-1 && r > l && nums[r] == nums[r+1] {
					r--
				}

			} else {
				if nums[l]+nums[r]+nums[i] < 0 {
					l++
				} else {
					r--
				}
			}
		}
	}
	return res
}
