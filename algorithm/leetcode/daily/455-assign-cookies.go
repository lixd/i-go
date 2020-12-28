package daily

import "sort"

// https://leetcode-cn.com/problems/assign-cookies/
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var i, j int
	for i < len(s) && j < len(g) {
		if s[i] >= g[j] {
			j++
		}
		i++
	}
	return j
}
