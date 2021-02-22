package lcof

/*
 2 <= n <= 1000 n取值比较大，所以不能用dp了，可能会越界。

当绳子长度大于4时，尽可能多的分成长度为3的小段，这样乘积是最大的。（数学证明自行查找）
所以只需要把n大于4时切分出3就行了
*/
func cuttingRope2(n int) int {
	// 2 <= n <= 1000
	if n <= 3 {
		return n - 1
	}
	var (
		sum int64 = 1
		mod       = int64(1000000007)
	)
	// 只要n大于4就切分一个3出来 等于4就别切了
	// 因为4最大切分为2*2还是等于4 切分为1*3反而变小了
	for n > 4 {
		sum *= 3
		sum %= mod
		n -= 3
	}
	sum *= int64(n)
	sum %= mod
	return int(sum)
}
