package daily

func characterReplacement(s string, k int) int {
	var (
		maxCnt, left int
		cnt          = [26]int{}
	)
	for right, ch := range s {
		// 每次遍历时记录当前区间内该字符总出现次数
		cnt[ch-'A']++
		// 找出最大值 然后 right-left+1 表示当前区间字符串 减去最大值就是剩下的不相等的字符串
		// 超过K则替换K个后也不能全部相等，所以需要移动left 进入下一个区间
		maxCnt = max(maxCnt, cnt[ch-'A'])
		if right-left+1-maxCnt > k {
			// 区间内不相等的字符大于k则当前区间无意义
			// left 往右移一位 然后把 cnt 中对应字符的计数减1
			cnt[s[left]-'A']--
			left++
		}
	}
	// 以上循环找到了满足条件的 left 的位置,left 右边的字符个数就是答案
	return len(s) - left
}
