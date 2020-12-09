package dfa

import (
	"fmt"
	"testing"
)

func TestDFA(t *testing.T) {
	dfa := NewDFA()

	dfa.Append([]rune("abc"))
	dfa.Append([]rune("abcdef"))
	contains := dfa.Contains("abc")
	fmt.Println(contains)
	search := dfa.Search("abc", MatchAll)
	fmt.Println(search)
	_, marked := dfa.Cover("abcff", '*')
	fmt.Println(marked)
}
