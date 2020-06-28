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
1 因为只有字母 所以使用字符数组代替hash表 性质都是一样的也可以使用hash表
2 设置两个索引坐标 也是滑动窗口的左右索引坐标
3 题中要求的不重复字符数组，当不存在重复的时候右坐标向右扩大一个位置
4 当存在重复的时候，左坐标向右扩大一个位置，直到找到窗口里重复的那个元素都需要被移除窗口

5 比较当前窗口里的字符串长度 获取最大
https://github.com/baxiang/leetcode-go
*/
func lengthOfLongestSubstring(s string) int {
	freq := make([]int, 128)
	var res = 0
	start, end := 0, -1
	for start < len(s) {
		// freq[s[end+1]] == 0  说明 数组中没有这个字符
		if end+1 < len(s) && freq[s[end+1]] == 0 {
			// end 右移 且 数组对应记录加1
			end++
			freq[s[end]]++
		} else {
			// 存在则说明出现了重复字符 start 右移 且 之前 start 位置的记录减1
			freq[s[start]]--
			start++
		}
		// 最后挑一个最大的 res
		if end-start+1 > res {
			res = end - start + 1
		}
	}
	return res
}
