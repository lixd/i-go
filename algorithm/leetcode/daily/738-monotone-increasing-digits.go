package daily

import "strconv"

// https://leetcode-cn.com/problems/monotone-increasing-digits/
func monotoneIncreasingDigits(n int) int {
	s := []byte(strconv.Itoa(n))
	i := 1
	// 从低位往高位 找到第一个不满足单调递增的位置 i
	for i < len(s) && s[i] >= s[i-1] {
		i++
	}
	if i < len(s) {
		// 从 i 开始往前把不是单调递增的前一位减少1 一直到满足单调递增
		for i > 0 && s[i] < s[i-1] {
			s[i-1]--
			i--
		}
		// 然后把后续全替换为9（因为前一位-1了）
		for i++; i < len(s); i++ {
			s[i] = '9'
		}
	}
	ans, _ := strconv.Atoi(string(s))
	return ans
}
