package main

import (
	"fmt"
)

/*
面试题 01.05. 一次编辑
字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
输入:
first = "pale"
second = "ple"
输出: True

输入:
first = "pales"
second = "pal"
输出: False
*/
func main() {
	first := "intention"
	second := "execution"
	fmt.Println(oneEditAway(first, second))
}

func oneEditAway(first string, second string) bool {
	flen := len(first)
	slen := len(second)
	diff := 0
	// 长度差大于1则肯定不行
	if flen-slen > 1 || slen-flen > 1 {
		return false
	}
	// 长度差为1或0 则挨个比较
	f, s, diff := 0, 0, 0
	for f < flen && s < slen {
		if first[f] != second[s] {
			diff++
			//长度不等则应该是错位的情况 把长的下标+1之后肯定会相等 如果不等则不满足条件
			if flen != slen {
				if flen > slen {
					f++
					continue
				}
				if slen > flen {
					s++
					continue
				}
			} else {
				// 长度相等则不移动 继续比较下一位
			}
		}
		f++
		s++
	}
	return diff <= 1
}
