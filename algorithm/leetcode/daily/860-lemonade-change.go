package daily

// https://leetcode-cn.com/problems/lemonade-change/
func lemonadeChange(bills []int) bool {
	var (
		five, ten int
	)
	for _, v := range bills {
		switch v {
		case 5:
			five++
		case 10:
			if five < 1 {
				return false
			}
			ten++
			five--
		case 20:
			// 3个5或者1个5 1个10 才能找零
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five > 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
