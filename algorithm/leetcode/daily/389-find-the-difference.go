package daily

// https://leetcode-cn.com/problems/find-the-difference/

func findTheDifference(s string, t string) byte {
	m := make(map[int32]int)
	for _, v := range s {
		m[v]++
	}
	for _, v := range t {
		m[v]--
		if m[v] < 0 {
			return byte(v)
		}
	}
	return 0
}
func findTheDifference2(s string, t string) byte {
	var sum int
	for _, v := range t {
		sum += int(v)
	}
	for _, v := range s {
		sum -= int(v)
	}
	return byte(sum)
}
