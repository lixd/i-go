package lcof

// https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/

// hash法 用哈希表存储各个元素，如果某个元素对应的位置已经有值了，说明该元素是重复元素
func findRepeatNumber(nums []int) int {
	var m = make(map[int]struct{}, len(nums))
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			return v
		} else {
			m[v] = struct{}{}
		}
	}
	return 0
}

// https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/solution/xue-sheng-wu-de-nu-peng-you-du-neng-kan-kdauw/
// 原地置换法 题目中有限制:在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。所以我们可以把 number i 放到第i个位置及，num[i]=i
// 如果出现重复元素比如出现以下情况：第一个元素A占据了A位置，后续把第二个A放到A位置的时候发现A位置已经存在一个A了，所以A就是重复元素
func findRepeatNumber2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		// 一直循环交换到位置i存放的是number i
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			// 不一致就交换位置i中的元素和位置i元素的正确位置中存放的元素。
			// 假设 i=1 num[1]中的数为2 那就需要把数字2放到位置2，所以进行交换num[1]和nums[2]
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return 0
}
