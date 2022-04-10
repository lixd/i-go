package number

func lengthOfLongestSubstring(s string) int {
	var m = make(map[byte]int)
	var left, right, ans int
	for left <= right && right < len(s) {
		if m[s[right]] != 0 { // 遇到重复时移动left，具体移动位置由map记录
			left = max(left, m[s[right]])
		}
		m[s[right]] = right + 1 // 记录下次遇到s[i]时left需要更新到的位置，所以是+1
		ans = max(ans, right-left+1)
		right++
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
