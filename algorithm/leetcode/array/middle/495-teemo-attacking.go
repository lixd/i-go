package middle

// findPoisonedDuration
// https://leetcode-cn.com/problems/teemo-attacking/
func findPoisonedDuration(timeSeries []int, duration int) int {
	var (
		res  int
		span int
	)
	if len(timeSeries) <= 1 {
		return len(timeSeries) * duration
	}
	for i := 1; i < len(timeSeries); i++ {
		span = timeSeries[i] - timeSeries[i-1]
		if span > duration {
			res += duration
		} else {
			res += span
		}
	}
	// 最后一次攻击必定持续 duration 时间
	res += duration
	return res
}
