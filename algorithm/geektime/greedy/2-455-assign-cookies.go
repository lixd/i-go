package greedy

import "sort"

// https://leetcode-cn.com/problems/assign-cookies/description/
func findContentChildren(g []int, s []int) int {
	// 排好序后一次遍历即可
	// 贪心算法 每个孩子都找刚好能满足其胃口的饼干从而使得饼干作用最大化
	sort.Ints(g)
	sort.Ints(s)
	var i, j int
	for i < len(g) && j < len(s) {
		// 如果找到满足条件的饼干则直接i++去匹配下一个孩子
		if g[i] <= s[j] {
			i++
		}
		// 不管满不满足都要j++，因为 g s都是有序的，当前不满足则只能往后找更大的饼干，满足则当前饼干被用掉了 也不能再用了
		j++
	}
	return i
}
