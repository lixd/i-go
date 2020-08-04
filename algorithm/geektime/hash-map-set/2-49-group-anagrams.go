package hash_map_set

// https://leetcode-cn.com/problems/group-anagrams/
// groupAnagrams 将数组作为 map的key value 为 []string
// 最后直接遍历整个 map 即可
func groupAnagrams(strs []string) [][]string {
	var (
		res [][]string
		m   = make(map[[26]int][]string)
	)
	if len(strs) == 0 {
		return nil
	}
	if len(strs) == 1 {
		return [][]string{strs}
	}

	for _, s := range strs {
		key := getKey(s)
		// 这里如果 计算出来的 key 相同则会自动存在一起
		m[key] = append(m[key], s)
	}
	// 然后遍历写入二维数组就是结果了
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// getKey 获取字符串 s 对应的数组 其实就是 askii 码值与 'a' 的差作为下标，出现一次则当前位置 +1
func getKey(s string) [26]int {
	arr := [26]int{}
	for _, v := range s {
		arr[v-'a']++
	}
	return arr
}
