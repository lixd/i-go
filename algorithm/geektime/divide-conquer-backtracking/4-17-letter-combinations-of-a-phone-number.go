package recursion

var (
	combinations []string
	phoneMap     map[string]string = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
)

// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

// backtrack 回溯
func backtrack(nums string, index int, combination string) {
	// terminator
	if index == len(nums) {
		// index == len(nums) 说明已经是走到最后了 拼装当前分支结果并返回
		combinations = append(combinations, combination)
		return
	}
	// 	process 否则取出当前层数字并找到对应字母，for循环出对应个数的分支并且index+1进入下一层
	num := string(nums[index])
	letters := phoneMap[num]
	for i := 0; i < len(letters); i++ {
		// drill down
		backtrack(nums, index+1, combination+string(letters[i]))
	}
	// 	reverse
}
