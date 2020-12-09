package dfa

import (
	"strings"
	"sync"
)

/*
DFA 敏感词过滤算法
基于Trie树实现
*/

type Trie interface {
	Append(string)
}

type Check interface {
	Trie
	Contains(string) bool
}

type Search interface {
	Trie
	Search(string) []Index
}
type Cover interface {
	Search
	Cover()
}

type Index struct {
	start int
	end   int
}

// DFA则是字典树上进行优化
type DFA struct {
	root *Node
	mu   sync.Mutex
}

const (
	MatchFirst = 0 // 匹配到第一个敏感词就结束
	MatchAll   = 1 // 匹配完整个字符串
)

// 字典树
type Node struct {
	children children
	isEnd    bool // 是否是单词的结束，如打人受伤  打(true) 人(true) 伤(true)，可以有三个词汇（打/打人/打人受伤）
}
type children map[rune]*Node // 每个子节点都可以包含多个节点

func NewDFA() *DFA {
	return &DFA{
		root: newTrieNode(),
	}
}

func newTrieNode() *Node {
	return &Node{
		children: make(children),
		isEnd:    false,
	}
}

func newDFANode() *Node {
	return &Node{
		children: make(children),
		isEnd:    false,
	}
}

// Append 增加敏感词库 构建Trie树
func (dfa *DFA) Append(words []rune) {
	if dfa.root == nil {
		return
	}

	dfa.mu.Lock()
	defer dfa.mu.Unlock()

	currNode := dfa.root
	for _, word := range words {
		targetNode, exist := currNode.children[word]
		if !exist {
			// 当前单词不存在则添加到Trie中
			targetNode = newDFANode()
			currNode.children[word] = targetNode
			currNode = targetNode
		} else {
			// 存在则移动到下一个节点继续判断
			currNode = targetNode
		}
	}
	// 最后一个节点是完整单词 修改标记
	currNode.isEnd = true
}

// Contains 校验输入字符串中是否存在敏感词
func (dfa *DFA) Contains(sentence string) bool {
	var (
		isMatchFull bool // 是否匹配到完整敏感词
		matchCount  int  // 匹配到单词个数
	)
	// 统一转换成小写
	sentenceRune := []rune(strings.ToLower(sentence))
	currNode := dfa.root
	for _, v := range sentenceRune {
		targetNode, exist := currNode.children[v]
		if exist {
			matchCount++
			currNode = targetNode
			if currNode.isEnd {
				isMatchFull = true
				break
			}
		} else {
			// 不匹配则跳转到下一个单词并还原到root节点 继续匹配
			currNode = dfa.root
		}
	}
	// 限制必须匹配到两个及其以上的敏感字才算（单个敏感字直接排除掉）
	if matchCount < 2 {
		return false
	}

	return isMatchFull
}

// Search 搜索给定字符串中的敏感字并返回出现的位置信息
func (dfa *DFA) Search(sentence string, matchType int64) (matchIndexList []Index) {
	sentenceRune := []rune(strings.ToLower(sentence))
	currNode := dfa.root
	var (
		start        int
		isFirstMatch = true // 是否是第一次匹配到敏感字
	)
	for i := 0; i < len(sentenceRune); i++ {
		targetNode, exist := currNode.children[sentenceRune[i]]
		if exist {
			if isFirstMatch {
				// 第一次匹配到则记录起点并将标记置为false
				start = i
				isFirstMatch = false
			}
			currNode = targetNode
			if currNode.isEnd {
				matchIndexList = append(matchIndexList, Index{start: start, end: i})
				// 如果只匹配第一个敏感词则直接返回
				if matchType == MatchFirst {
					return matchIndexList
				}

				// 重新回到root节点并重置标记，继续寻找下一个敏感词
				currNode = dfa.root
				isFirstMatch = true
				start = -1
			}
		} else {
			// 如果匹配到了敏感字则下一轮从第一个敏感字的后一个字开始匹配
			if !isFirstMatch {
				// 应该把i修改为start+1 但是本轮循环结束后会执行i++
				// 所以这里就赋值为start
				i = start
			}
			// 重新回到root节点并重置标记，继续寻找下一个敏感词
			currNode = dfa.root
			start = -1
			isFirstMatch = true
		}
	}

	return matchIndexList
}

// Cover 找出给定字符串中的铭感词并将其替换为指定字符{mask}
func (dfa *DFA) Cover(sentence string, mask rune) (ok bool, marked string) {
	matchIndexList := dfa.Search(sentence, MatchAll)
	if len(matchIndexList) == 0 {
		return false, sentence
	}

	sentenceRune := []rune(sentence)
	for _, matchIndexStruct := range matchIndexList {
		for i := matchIndexStruct.start; i <= matchIndexStruct.end; i++ {
			sentenceRune[i] = mask
		}
	}

	return true, string(sentenceRune)
}
