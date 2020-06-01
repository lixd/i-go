package main

import (
	"fmt"
	"strings"
)

/*
面试题 01.09. 字符串轮转
字符串轮转。给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成（比如，waterbottle是erbottlewat旋转后的字符串）
 输入：s1 = "waterbottle", s2 = "erbottlewat"
 输出：True
旋转体的两倍必然包含另一个旋转体
*/
func main() {
	s1 := "waterbottle"
	s2 := "erbottlewat"
	fmt.Println(isFlipedString(s1, s2))
}
func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	return strings.Contains(s2+s2, s1)
}
