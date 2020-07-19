package array_linkedlist_skiplist

// https://leetcode-cn.com/problems/two-sum
// 将 数组的值 和 下标 交换后存到 map 中 map 的 key 为 值 value 为下标
func twoSum(nums []int, target int) []int {
	var m = make(map[int]int)
	for i, v := range nums {
		// 找出与当前值相加等于 target 的目标值
		find := target - v
		// 由于是反着存的 所以 如果 m[find]如果存在 那么他的 value 就是 数组的 index
		if j, ok := m[find]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return []int{}
}
