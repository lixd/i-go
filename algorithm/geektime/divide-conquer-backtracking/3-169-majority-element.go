package recursion

// https://leetcode-cn.com/problems/majority-element/
// 题解 https://leetcode-cn.com/problems/majority-element/
// 众数是出现次数大于二分之一的数 因此按照抵消法抵消后 最后剩余的一定是众数
/*
1.利用map统计元素出现的次数
2.利用大于 ⌊ n/2 ⌋ 的特性 该方法就是用到了这个特性 摩尔投票算法
3.先排序，在查找
4.暴力破解，双重for循环就不写了
*/
func majorityElement(nums []int) int {
	var (
		major = 0
		count = 0
	)
	for _, v := range nums {
		if count == 0 {
			major = v
		}
		if major == v {
			count++
		} else {
			count--
		}

	}
	return major
}
