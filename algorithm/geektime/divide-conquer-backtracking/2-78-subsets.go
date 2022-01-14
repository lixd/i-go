package recursion

// https://leetcode-cn.com/problems/subsets/
// 递归
func subsets(nums []int) [][]int {
	var result [][]int
	if len(nums) == 0 {
		return result
	}
	dfs(&result, nums, []int{}, 0)
	return result
}

func dfs(result *[][]int, nums []int, curr []int, index int) {
	// 1.terminator
	if index == len(nums) {
		// 这里需要 copy一个出来
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*result = append(*result, tmp)
		return
	}
	// 2.process logic
	// 3.drill down
	// 每层都可以选择加当前数 或者不加当前数
	dfs(result, nums, curr, index+1)
	curr = append(curr, nums[index])
	dfs(result, nums, curr, index+1)
	// 4.reverse states
	// 最后需要把list复原，防止干扰到其他层级
	curr = curr[:len(curr)-1]
}

// 迭代
func subsets2(nums []int) [][]int {
	result := [][]int{[]int{}}
	for i := 0; i < len(nums); i++ {
		for _, old := range result {
			// cpoy出一个和old数组 len cap 都一样的数组
			new := make([]int, len(old), cap(old)+1)
			copy(new, old)
			// 添加当前数值到数组中
			new = append(new, nums[i])
			result = append(result, new)
		}
	}
	return result
}
