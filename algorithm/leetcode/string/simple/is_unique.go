package main

import (
	"fmt"
)

/*
面试题 01.01. 判定字符是否唯一
实现一个算法，确定一个字符串 s 的所有字符是否全都不同。

示例 1：

输入: s = "leetcode"
输出: false
示例 2：

输入: s = "abc"
输出: true
限制：

0 <= len(s) <= 100
如果你不使用额外的数据结构，会很加分。

*/
func main() {
	astr := "leetcode"
	fmt.Println(isUnique(astr))
}

func isUnique(astr string) bool {
	var m = make(map[int32]bool)
	arr := []rune(astr)
	for i := 0; i < len(arr); i++ {
		if _, ok := m[arr[i]]; ok {
			return false
		}
		m[arr[i]] = true
	}
	return true
}
