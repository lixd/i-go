package daily

// https://leetcode-cn.com/problems/binary-prefix-divisible-by-5/
/*
判定能否被5整除 只需要看最后一位是不是0或者5 所以直接模10获取最后一位 这样还可以防止溢出
然后每次都要在末尾加一个数 所以直接左移然后模10最后加上当前数即可
*/
func prefixesDivBy5(A []int) []bool {
	ans := make([]bool, len(A))
	// r
	r := 0
	for i, n := range A {
		// 01<<1 + 1 =011
		r = (r<<1)%10 + n
		ans[i] = r%5 == 0
	}

	return ans
}
