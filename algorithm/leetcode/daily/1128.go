package daily

// https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs/
func numEquivDominoPairs(dominoes [][]int) int {
	var (
		cnt = [100]int{} // 用哈希表来记录对应key出现的次数，由于只有100个值 所以优化成数组
		ans int
	)
	for _, d := range dominoes {
		// 默认用小的值*10+大的值 作为数组的索引
		if d[0] > d[1] {
			// 如果d[0]为大值则和小值d[1]交换
			// 这样[1,2] [2,1] 最终交换后都是[1,2]这样的形式
			d[0], d[1] = d[1], d[0]
		}
		i := d[0]*10 + d[1] // 计算对应的索引 可以当做是一个简单的hash函数
		ans += cnt[i]
		cnt[i]++
	}
	return ans
}
