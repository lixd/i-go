package daily

// https://leetcode-cn.com/problems/add-to-array-form-of-integer/
/*
逐位相加模板
while ( A 没完 || B 没完))
    A 的当前位
    B 的当前位

    和 = A 的当前位 + B 的当前位 + 进位carry

    当前位 = 和 % 10;
    进位 = 和 / 10;

判断还有进位吗
*/
func addToArrayForm(A []int, K int) []int {
	var (
		carry int
		i     = len(A) - 1
		ret   = make([]int, 0, len(A))
	)

	for i >= 0 || K != 0 {
		a, b := 0, 0
		if i >= 0 {
			a = A[i]
		}
		if K != 0 {
			b = K % 10
		}
		sum := a + b + carry
		carry = sum / 10
		K = K / 10
		i--
		ret = append(ret, sum%10)
	}
	if carry != 0 {
		ret = append(ret, carry)
	}
	reverse(ret)
	return ret
}

func reverse(a []int) {
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
