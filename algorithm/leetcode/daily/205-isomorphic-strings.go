package daily

// https://leetcode-cn.com/problems/isomorphic-strings
func isIsomorphic(s string, t string) bool {
	var (
		s2t = make(map[byte]byte)
		t2s = make(map[byte]byte)
	)
	for i := range s {
		x := s[i]
		y := t[i]
		// 分别检查两个map 其中一个有冲突则说明不是同构字符串
		if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
			return false
		}
		//交叉存 方便后续遇到同样的字符串时好比较
		s2t[x] = y
		t2s[y] = x
	}
	return true
}
