package main

import (
	"fmt"
	"strings"
)

/*
面试题 01.06. 字符串压缩
字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。
若“压缩”后的字符串没有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。

*/
func main() {
	S := "aabcccccaa"
	fmt.Println(compressString(S))
}

func compressString(S string) string {
	// 长度小于等于1直接返回
	if len(S) <= 1 {
		return S
	}

	var sb strings.Builder
	prev := S[0]
	count := 1

	for i := 1; i < len(S); i++ {
		if prev == S[i] {
			count++
		} else {
			sb.WriteByte(prev)
			sb.WriteString(fmt.Sprintf("%d", count))
			prev = S[i]
			count = 1
		}

		// 纪录最后结尾的字符
		if i == len(S)-1 {
			sb.WriteByte(prev)
			sb.WriteString(fmt.Sprintf("%d", count))
		}
	}
	if sb.Len() < len(S) {
		return sb.String()
	}
	return S
}
