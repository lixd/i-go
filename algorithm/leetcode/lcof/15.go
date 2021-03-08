package lcof

// https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
func hammingWeight(num uint32) int {
	var sum int
	for num != 0 {
		// 只有1&1=1
		// 1的位数和num不一致，默认会在前面补0，所以这些补出来的位数&只会结果都为0 最终影响结果的就只有num的最后一位
		// 111x & 1 --> 111x & 0001 --> x & 1
		sum += int(num & 1)
		num >>= 1
	}
	return sum
}
