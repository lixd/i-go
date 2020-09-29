package greedy

// https://leetcode-cn.com/problems/lemonade-change/description/
func lemonadeChange(bills []int) bool {
	var five, ten int
	for _, bill := range bills {
		switch bill {
		case 5:
			five++
		case 10:
			if five < 1 {
				return false
			}
			ten++
			five--
		case 20:
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five > 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
