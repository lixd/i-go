package main

import (
	"bytes"
	"fmt"
	"strings"
)

/*
面试题 01.03. URL化
URL化-替换空格。编写一种方法，将字符串中的空格全部替换为%20。假定该字符串尾部有足够的空间存放新增字符，并且知道字符串的“真实”长度。（注：用Java实现的话，请使用字符数组实现，以便直接在数组上操作。）

示例1:

 输入："Mr John Smith    ", 13
 输出："Mr%20John%20Smith"
示例2:

 输入："               ", 5
 输出："%20%20%20%20%20"
提示：

字符串长度在[0, 500000]范围内。
*/
func main() {
	S := "               "
	length := 5
	fmt.Println(replaceSpaces(S, length))
}

// 36ms
func replaceSpaces(S string, length int) string {
	var (
		result = make([]int32, 0, len(S))
	)
	for i := 0; i < length; i++ {
		if string(S[i]) == " " {
			result = append(result, '%')
			result = append(result, '2')
			result = append(result, '0')
		} else {
			result = append(result, int32(S[i]))
		}
	}

	return string(result)
}

// 32ms
func replaceSpaces2(S string, length int) string {
	var buffer bytes.Buffer
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			buffer.WriteString("%20")
		} else {
			buffer.WriteString(string(S[i]))
		}
	}
	return buffer.String()
}

// 24ms
func replaceSpaces3(S string, length int) string {
	return strings.ReplaceAll(S[:length], " ", "%20")
}

// 16ms
func replaceSpaces4(S string, length int) string {
	num := 0
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			num++
		}
	}
	result := make([]byte, 3*num+(length-num))
	k := 0
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			result[k] = '%'
			result[k+1] = '2'
			result[k+2] = '0'
			k += 3
		} else {
			result[k] = S[i]
			k++
		}
	}
	return string(result)
}
