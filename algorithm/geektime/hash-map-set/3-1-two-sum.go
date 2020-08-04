package hash_map_set

// https://leetcode-cn.com/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
	var m = make(map[int]int)
	for i, v := range nums {
		find := target - v
		if k, ok := m[find]; ok {
			return []int{k, i}
		}
		m[v] = i
	}
	return nil
}
