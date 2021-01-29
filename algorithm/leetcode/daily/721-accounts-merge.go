package daily

import "sort"

// https://leetcode-cn.com/problems/accounts-merge/
// 并查集 https://zhuanlan.zhihu.com/p/93647900/
func accountsMerge(accounts [][]string) (ans [][]string) {
	emailToIndex := map[string]int{}
	emailToName := map[string]string{}
	for _, account := range accounts {
		name := account[0]
		for _, email := range account[1:] {
			if _, has := emailToIndex[email]; !has {
				emailToIndex[email] = len(emailToIndex)
				emailToName[email] = name
			}
		}
	}

	parent := make([]int, len(emailToIndex))
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(from, to int) {
		parent[find(from)] = find(to)
	}

	for _, account := range accounts {
		firstIndex := emailToIndex[account[1]]
		for _, email := range account[2:] {
			union(emailToIndex[email], firstIndex)
		}
	}

	indexToEmails := map[int][]string{}
	for email, index := range emailToIndex {
		index = find(index)
		indexToEmails[index] = append(indexToEmails[index], email)
	}

	for _, emails := range indexToEmails {
		sort.Strings(emails)
		account := append([]string{emailToName[emails[0]]}, emails...)
		ans = append(ans, account)
	}
	return
}
