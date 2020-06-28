package main

import "fmt"

/*
寻找两个正序数组的中位数 困难

给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
你可以假设 nums1 和 nums2 不会同时为空。

nums1 = [1, 3]
nums2 = [2]
则中位数是 2.0

nums1 = [1, 2]
nums2 = [3, 4]
则中位数是 (2 + 3)/2 = 2.5
*/
func main() {
	nums1 := []int{}
	nums2 := []int{}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

}
