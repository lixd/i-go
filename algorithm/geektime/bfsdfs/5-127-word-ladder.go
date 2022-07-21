package bfsdfs

// https://leetcode-cn.com/problems/word-ladder/description/
// ladderLength
/*
具体逻辑为根据当前word去wordList中寻找与之只差一个的字符串（因为规定了每次只能改变一个）
如果有多个则下一层级由这多个字符串继续发散出去 继续再wordList中寻找与之只差一个的字符串
直到出现endWord返回
需要记录过程中哪个word已经被使用了 否则会出现死循环
*/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 不存在直接返回
	if !isContainer(endWord, wordList) {
		return 0
	}

	step := 0
	// used 用于存储wordList中哪个已经被使用过了 否则会出现死循环 即AA-->AB-->AA
	used := make([]bool, len(wordList))
	// queue 则存储当前层级
	queue := []string{beginWord}
	for len(queue) > 0 {
		step++
		// next 存储下一层级
		next := make([]string, 0)
		for _, v := range queue {
			if v == endWord {
				return step
			}
			for j, w := range wordList {
				if !used[j] && hasOneDiff(v, w) {
					next = append(next, w)
					used[j] = true
				}
			}
		}
		queue = next
	}
	return 0
}

// isContainer 检查 str 是否存在于 bank 中
func isContainer(str string, bank []string) bool {
	for _, s := range bank {
		if str == s {
			return true
		}
	}
	return false
}

// hasOneDiff 判断 x,y 两个字符串是否只差一个字符
func hasOneDiff(x, y string) bool {
	count := 0
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			count++
		}
		if count > 1 {
			return false
		}
	}
	return count == 1
}
