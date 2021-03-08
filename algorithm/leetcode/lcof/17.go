package lcof

import (
	"math"
	"strings"
)

// https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/
// 简单解法 不考虑大数问题
func printNumbers(n int) []int {
	var (
		limit = math.Pow(10, float64(n))
		ans   = make([]int, 0, int(limit-1))
	)
	for i := 1; i < int(limit); i++ {
		ans = append(ans, i)
	}
	return ans
}

/*
首先要考虑到当n很大的时候（比如100），打印出来的数很有可能是超过了INT_MAX的范围的，所以我们用字符串来表示每个数。
每一位有0~9 这10个值可以选 所以整个结果是一个树形结构
第一位0~9中选一个 第二位又是0~9中选 直接使用递归
*/
func printNumbersOverflow(n int) []string {
	var (
		limit = math.Pow(10, float64(n))
		ans   = make([]string, 0, int(limit-1))
	)
	numberDFS(0, n, "", &ans)
	return ans
}

var s = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func numberDFS(level, n int, str string, ans *[]string) {
	// terminator
	if level == n {
		if str != strings.Repeat("0", n) {
			str = strings.TrimPrefix(str, "0")
			*ans = append(*ans, str)
		}
		return
	}
	// drill down
	for _, v := range s {
		numberDFS(level+1, n, str+v, ans)
	}
}
