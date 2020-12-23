package daily

// https://leetcode-cn.com/problems/first-unique-character-in-a-string/
func firstUniqChar(s string) int {
	ch := make([]int, 26)
	for _, v := range s {
		ch[v-'a']++
	}
	for i, v := range s {
		c := ch[v-'a']
		if c == 1 {
			return i
		}
	}
	return -1
}
