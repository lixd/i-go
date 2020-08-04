package hash_map_set

// https://leetcode-cn.com/problems/valid-anagram/description/
// 1.阅读题目并且沟通不清楚的地方
// 2.想出可能的解法并找到最优解
// 	暴力解法、sort 排序后比较、hash map记录字符出现次数
// 3.code
// 4.test
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if s == t {
		return true
	}
	arr := make([]int, 26)
	for _, v := range s {
		// v-'a' 计算出 askii 码差距最大为 26 所以这样只需要一个大小为 26 的数组
		// 直接用 arr[v]++ 则需要更大的数组
		arr[v-'a']++
	}
	for _, v := range t {
		arr[v-'a']--
	}
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}
