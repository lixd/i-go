package lcof

// https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/
/*
最简单的办法自然是for循环x*x*x 一个乘就好了 这样时间复杂度为O(n)
使用快速幂方法能降低为O(logn)
x^16=x^8*x^2
x^8=x^4*x^2
x^4=x^2*x^2
求出x^2之后就没必要在重复求了，所以这样能降低时间复杂度
*/
func myPow(x float64, n int) float64 {
	if n < 0 {
		// n为负数那么结果就是就是倒数 x^-2 --> 1/x^2
		return 1.0 / pow(x, -n)
	}
	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	sub := pow(x, n/2) // 利用分治思想 将n拆分下去逐个计算
	sub *= sub         // n/2 这里缩小了两次方 所以需要补上
	// 区分奇数还是偶数 因为 n/2 会有误差 比如 5/2=2 如果n是奇数则会漏掉一个 也要补上
	if n%2 != 0 {
		sub *= x
	}
	return sub
}
