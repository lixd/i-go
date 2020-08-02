package array_linkedlist_skiplist

// https://leetcode-cn.com/problems/merge-sorted-array
// 从后往前写 不需要额外空间
func merge(nums1 []int, m int, nums2 []int, n int) {
	var (
		pm = m - 1
		pn = n - 1
		p  = m + n - 1
	)
	for pn >= 0 {
		if pm >= 0 && nums1[pm] > nums2[pn] {
			nums1[p] = nums1[pm]
			pm--
			p--
		} else {
			nums1[p] = nums2[pn]
			pn--
			p--
		}
	}
}
