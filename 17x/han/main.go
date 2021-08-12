package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func main() {
	s1 := "我是中国人hello word!,2020 street 188#"
	var count int
	for _, v := range s1 {
		if unicode.Is(unicode.Han, v) {
			fmt.Println("找到中文")
			count++
		}
	}
	fmt.Println(count)
	fmt.Println(IsChineseChar(s1))

}

// IsChineseChar 判断字符串是否包含中文
func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}
