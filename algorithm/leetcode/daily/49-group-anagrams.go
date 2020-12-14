package daily

import "sort"

// https://leetcode-cn.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	// 互为字母异位词的字符串是相同的，只是顺序不一致，所以先进行排序 排序后肯定相同
	// 把排序后的值作为 key 即可
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ret := make([][]string, 0, len(mp))
	for _, v := range mp {
		ret = append(ret, v)
	}
	return ret
}
