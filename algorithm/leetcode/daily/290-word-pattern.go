package daily

import "strings"

// https://leetcode-cn.com/problems/word-pattern/
func wordPattern(pattern string, s string) bool {
	// 长度不一致直接返回
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	// 用两个map分别存储pattern-->word 和 word-->pattern 的对应关系
	// key:pattern value word
	char2word := make(map[int32]string)
	// key:word value:pattern
	word2char := make(map[string]int32)
	for i, v := range pattern {
		// 二者都为空则存入新的对应关系
		_, okPw := char2word[v]
		char, okWp := word2char[words[i]]
		if !okPw && !okWp {
			char2word[v] = words[i]
			word2char[words[i]] = v
			continue
		}
		// 否则 可能其中一个已经用过了
		// 通过 char 是否一致来判断本次是重复使用还是冲突
		if v != char {
			return false
		}
	}
	return true
}

func wordPattern2(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	char2word := make(map[int32]string)
	word2char := make(map[string]int32)
	for i, v := range pattern {
		word := words[i]
		// word2char 或者 char2word 任意一个冲突都说明不满足条件
		if word2char[word] > 0 && word2char[word] != v || char2word[v] != "" && char2word[v] != word {
			return false
		}
		char2word[v] = words[i]
		word2char[words[i]] = v
	}
	return true
}
