package main

import "fmt"

/*
无重复字符的最长子串 中等
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/
func main() {
	s := "xx"
	fmt.Println(lengthOfLongestSubstring(s))
}

// lengthOfLongestSubstring 双指针之滑动窗口
/*
我们可以使用「滑动窗口」来解决这个问题:
1) 我们使用两个指针表示字符串中的某个子串（的左右边界）。其中左指针代表着上文中「枚举子串的起始位置」，而右指针即为结束位置
2) 在每一步的操作中，我们会将左指针向右移动一格，表示 我们开始枚举下一个字符作为起始位置，然后我们可以不断地向右移动右指针，但需要保证这两个指针对应的子串中没有重复的字符。在移动结束后，这个子串就对应着 以左指针开始的，不包含重复字符的最长子串。我们记录下这个子串的长度；
3) 在枚举结束后，我们找到的最长的子串的长度即为答案。
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	// window 用于计数
	window := map[byte]int{}
	// 左右指针默认都从 0 开始
	left, right := 0, 0
	// length 为结果
	length := 0
	for right < len(s) {
		// 每次过一个字符则 right 加1
		c := s[right]
		right++
		// 存在 map 中 以字符 为 key
		window[c]++
		if left == 0 {
			length = right - 1
		}
		// value 不等于 1 则说明 遇到了两个重复的字符
		for window[c] != 1 {
			d := s[left]
			left++
			window[d]--
		}
		if right-left > length {
			length = right - left
		}
	}
	return length
}

/*
1 因为只有字母 所以使用字符数组代替hash表 性质都是一样的也可以使用hash表
2 设置两个索引坐标 也是滑动窗口的左右索引坐标
3 题中要求的不重复字符数组，当不存在重复的时候右坐标向右扩大一个位置
4 当存在重复的时候，左坐标向右扩大一个位置，直到找到窗口里重复的那个元素都需要被移除窗口

5 比较当前窗口里的字符串长度 获取最大
https://github.com/baxiang/leetcode-go
作者：ba-xiang
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/go-hua-dong-chuang-kou-by-ba-xiang/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func lengthOfLongestSubstring2(s string) int {
	freq := make([]int, 128)
	var res = 0
	start, end := 0, -1
	for start < len(s) {
		if end+1 < len(s) && freq[s[end+1]] == 0 {
			end++
			freq[s[end]]++
		} else {
			freq[s[start]]--
			start++
		}
		res = max(res, end-start+1)
	}
	return res
}
func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
