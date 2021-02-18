package lcof

// https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
/*
类似 斐波那契数列
区别是前面几阶:
f(0)=1 默认开始就在第0阶，所以也算是一种
f(1)=1 第一阶肯定只有一种
f(2)=2 可以从0阶一次走两阶上来，也可以从1阶一次走1阶走上来
第n阶台阶可以从第n-1阶一次走一阶走上去，也可以从n-2阶一次走两阶走上去
所以:fn=f(n-1)+f(n-2)
*/
func numWays(n int) int {
	var mod = 1000000007
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	var f, n1, n2 = 2, 1, 1
	for i := 3; i <= n; i++ {
		n2 = n1
		n1 = f
		f = (n1 + n2) % mod
	}
	return f
}
