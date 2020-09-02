package recursion

// https://leetcode-cn.com/problems/powx-n/
// 1.暴力解法 for 循环 把x乘N次
// 2.分治 1.terminator 2.process（spilt  problem） 3.drill down（subproblem）,merge result 4.reverse states
// 		pow（x,n）-->pow(x,n/2) n需要区分正负和奇偶
// 3.牛顿迭代法
func myPow(x float64, n int) float64 {
	if n < 0 {
		// 负数和正数区分
		return 1.0 / pow(x, -n)
	}
	return pow(x, n)
}
func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	sub := pow(x, n/2)
	// odd even 区分
	if n%2 == 0 {
		return sub * sub
	}
	// odd 除2之后会漏掉一个 所以最后补上
	return sub * sub * x
}
