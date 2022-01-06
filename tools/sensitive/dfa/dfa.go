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

// DFA 则是字典树上进行优化
type DFA struct {
	root *Node
	mu   sync.Mutex
}

const (
	MatchFirst = 0 // 匹配到第一个敏感词就结束
	MatchAll   = 1 // 匹配完整个字符串
)

// Node 字典树
type Node struct {
	children map[rune]*Node
	isEnd    bool // 是否是单词的结束，如打人受伤  打(true) 人(true) 受(false)伤(true)，可以有三个词汇（打/打人/打人受伤）
}

func NewDFA() *DFA {
	return &DFA{
		root: newTrieNode(),
	}
}

func newTrieNode() *Node {
	return &Node{
		children: make(map[rune]*Node),
		isEnd:    false,
	}
}

// Append 增加敏感词库 构建Trie树
func (dfa *DFA) Append(words []rune) {
	if dfa.root == nil {
		return
	}
	if len(words) < 2 { // 限制不能添加单个字符的敏感词
		return
	}

	dfa.mu.Lock()
	defer dfa.mu.Unlock()
	// 挨个匹配，找到对应的位置添加进去
	// 第一个字符匹配第一层节点，第二个字符则匹配第二层节点，以此类推
	currNode := dfa.root // 默认从主节点开始
	for _, word := range words {
		targetNode, exist := currNode.children[word]
		if exist { // 存在则移动到下一个节点，继续判断下一个字符是否存在于下一个节点上。
			currNode = targetNode
		} else { // 不存在则新增一个节点，将当前字符添加到Trie树中
			targetNode = newTrieNode()
			currNode.children[word] = targetNode
			currNode = targetNode // 这里同样要把当前节点移动到新增节点，让后续字符都在这个节点的子节点上进行判断
		}
	}
	// 循环结束后，修改当前节点标记
	// NOTE: 不是真正意义上的完整单词，只能表明使用者往Trie树中添加过这个单词。
	// 比如输入 abc,那么 c 节点会被标记成完成单词
	currNode.isEnd = true
}

// Contains 校验输入字符串中是否存在敏感词
func (dfa *DFA) Contains(sentence string) bool {
	// 统一转换成小写
	sentenceRune := []rune(strings.ToLower(sentence))
	for i := 0; i < len(sentenceRune); i++ {
		currNode := dfa.root
		for j := i; j < len(sentenceRune); j++ {
			word := sentenceRune[j]
			targetNode, exist := currNode.children[word]
			if exist {
				currNode = targetNode
				if currNode.isEnd {
					return true
				}
			} else { // 不匹配则退出本轮循环，开始下一轮循环
				// 比如输入abcd,此时Trie树中只有abd、bcd,第一轮匹配ab成功，到c匹配失败，此时需要返回root节点从b开始第二轮匹配，而不是直接从d开始匹配。
				break
			}
		}
	}
	return false
}

// HasPrefix 是否存在以 prefix 为前缀的词则返回 true 否则返回 false
func (dfa *DFA) HasPrefix(prefix string) bool {
	// 统一转换成小写
	sentenceRune := []rune(strings.ToLower(prefix))
	currNode := dfa.root
	for _, v := range sentenceRune {
		targetNode, exist := currNode.children[v]
		if exist {
			currNode = targetNode
		} else { // 某个节点不存在则说明不存在这个前缀
			return false
		}
	}
	// 能走到最后说明都存在
	return true
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

// Cover 找出给定字符串中的敏感词并将其替换为指定字符 {mask}
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
