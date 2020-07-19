package array_linkedlist_skiplist

// https://leetcode-cn.com/problems/move-zeroes/
func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
}

// snowball 解法 https://leetcode.com/problems/move-zeroes/discuss/172432/THE-EASIEST-but-UNUSUAL-snowball-JAVA-solution-BEATS-100-(O(n))-%2B-clear-explanation
// 通过雪球个数来记录当前位置之前有多少个 0,每次遇到 0 则 snowBallSize 加 1.
// 交换时则把当前位置（index）和第一个 0（index-snowBallSize） 进行交换。
// 上面解法中通过 j 来记录第一个 0 的位置，这里则是用当前位置减去雪球个数来得到 第一个 0 的位置。
func moveZeroes2(nums []int) {
	snowBallSize := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			snowBallSize++
		} else if nums[i] != 0 {
			nums[i-snowBallSize], nums[i] = nums[i], nums[i-snowBallSize]
		}
	}
}

// moveZeroes 遍历数组 当前位置是 0 则不处理，否则将当前位置和 第一个 0 的位置互换
func moveZeroes3(nums []int) {
	var zero int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[zero] = nums[zero], nums[i]
			zero++
		}
	}
}
