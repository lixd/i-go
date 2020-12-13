package daily

//https://leetcode-cn.com/problems/contains-duplicate/
func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			return true
		} else {
			m[v] = struct{}{}
		}
	}
	return false
}
