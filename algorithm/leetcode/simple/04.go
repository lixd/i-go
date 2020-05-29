package main

import "fmt"

/*
面试题 01.04. 回文排列
给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。

回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。

回文串不一定是字典当中的单词。



示例1：

输入："tactcoa"
输出：true（排列有"tacocat"、"atcocta"，等等）

*/
func main() {
	s := "acv"
	fmt.Println(canPermutePalindrome(s))
}
func canPermutePalindrome(s string) bool {
	var (
		m     = make(map[byte]int)
		count int
	)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for _, v := range m {
		if v%2 == 1 {
			count++
		}
		if count > 1 {
			return false
		}
	}

	return true
}
