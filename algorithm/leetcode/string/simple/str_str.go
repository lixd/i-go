package main

import "fmt"

/*
实现 strStr() 函数。 简单

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

输入: haystack = "hello", needle = "ll"
输出: 2

输入: haystack = "aaaaa", needle = "bba"
输出: -1

*/
func main() {
	haystack := "hello world"
	needle := "world"
	fmt.Printf("strStr:%v \n", strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	var i, j int
	for i = 0; i < len(haystack)-len(needle)+1; i++ {
		for j = 0; j < len(needle); j++ {
			// 进入内层后 i 不移动 所以这里是 haystack[i+j]
			if haystack[i+j] != needle[j] {
				break
			}
		}
		// 每次内层循环退出后 判断一次 长度是否相等
		// 相等则说明找到了
		if len(needle) == j {
			return i
		}
	}
	return -1
}
