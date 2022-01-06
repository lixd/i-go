package dfa

import (
	"fmt"
	"testing"
)

func TestDFA(t *testing.T) {
	dfa := NewDFA()

	dfa.Append([]rune("abd"))
	dfa.Append([]rune("bcd"))
	hasPrefix := dfa.HasPrefix("abc")
	fmt.Println("是否存在前缀:", hasPrefix)
	contains := dfa.Contains("abcd")
	fmt.Println("是否存在敏感词:", contains)
	search := dfa.Search("abcdef", MatchAll)
	fmt.Println("敏感词出现位置:", search)
	_, marked := dfa.Cover("abcdef", '*')
	fmt.Println("对敏感词进行标记:", marked)
}
