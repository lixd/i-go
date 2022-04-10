package leetcode

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{i, j}
		}
		m[num] = i
	}
	return nil
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		head, tail *ListNode
		carry      int
	)
	for l1 != nil || l2 != nil {
		var c int
		// 	获取l1 l2需要先判空 获取后要将其指向next
		if l1 != nil {
			c += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			c += l2.Val
			l2 = l2.Next
		}
		// 先判断进位值
		if carry != 0 {
			c += carry
			carry = 0
		}
		// 累计值大于等于10就需要进位
		if c >= 10 {
			carry = c / 10
			c = c % 10
		}
		// 如果没有head则创建，因为需要返回head，所以单独哪个tail来记录链表尾部
		if head == nil {
			head = &ListNode{Val: c}
			tail = head
		} else {
			tail.Next = &ListNode{Val: c}
			tail = tail.Next
		}
	}
	// 到最后都还有进位值则再添加一个
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		head, tail *ListNode
		carry      int
	)
	for l1 != nil || l2 != nil {
		var sum int
		// 	获取l1 l2需要先判空 获取后要将其指向next
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carry
		// 进位这部分可以提出来，不需要单独判断
		carry = sum / 10
		sum %= 10
		// 如果没有head则创建，因为需要返回head，所以单独哪个tail来记录链表尾部
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	// 到最后都还有进位值则再添加一个
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	left, right, ans := 0, 0, 0
	m := map[byte]int{}
	for right < len(s) {
		_, ok := m[s[right]]
		if !ok {
			m[s[right]] = right
		} else {
			fmt.Printf("m[s[right]]+1 :%v left:%v\n", m[s[right]]+1, left)
			if m[s[right]]+1 >= left {
				left = m[s[right]] + 1
			}
			m[s[right]] = right
		}
		ans = max(right-left+1, ans)
		right++
	}

	return ans

}

func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}

	var left, right, ans int
	var m = make(map[byte]int)
	for right < len(s) {
		_, ok := m[s[right]]
		if !ok { // 不存在则写入map中
			m[s[right]] = right
		} else {
			// 	存在则需要移动left，直接移动到上一个重复字符出现的地方的+1,比如pwwkew，在第三个字符即w的时候发现重复了，那就直接left移动到所以2这里，
			// 需要注意left不能往左移，要先判断一下大小，由于没有把之前出现的字符从map中移除，所以可能出现比left小的情况
			if m[s[right]]+1 >= left {
				left = m[s[right]] + 1
			}
			m[s[right]] = right // 同时还是需要记录right
		}
		ans = max(ans, right-left+1)
		right++
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	begin, maxLen := 0, 1 // maxLen最少都是1，因为一个字符肯定是回文
	// 竖着填表 所以j固定先把i枚举完
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 { // 相差小于3就表示中间没有多余的子串了,相差为2中间就只有1个字符串，为1中间就没有字符串了，二者都肯定是回文
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var b = []rune(s)
	var res = make([]string, numRows)
	var length = len(b)
	var period = numRows*2 - 2 // 2n-2为一个循环
	for i := 0; i < length; i++ {
		mod := i % period  // 和每轮循环次数取模 去掉多余的部分
		if mod < numRows { // 小于行数则处理竖部分 直接写入即可
			res[mod] += string(b[i])
		} else { // 大于或者等于就算是在斜边部分了，
			// 比如row为4，每轮循环就是6，前面的0123都是一一对应的，只有后续的45在斜边上，需要特殊处理
			// 具体公式就是 period-mod,比如4就对应到2,5就是1,6则是下一个循环的0了.
			res[period-mod] += string(b[i])
		}
	}
	return strings.Join(res, "")
}

func convert2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	res := make([]string, numRows)
	loop := numRows*2 - 2
	for i := 0; i < len(s); i++ {
		mod := i % loop
		if mod < numRows {
			res[mod] += string(s[i])
		} else {
			res[loop-mod] += string(s[i])
		}

	}
	return strings.Join(res, "")
}

func reverse(x int) int {
	var ans int
	for x != 0 {
		mod := x % 10
		x /= 10
		// 如果ans本轮扩大只会超过int32就返回0
		if ans > 0 && ans > (math.MaxInt32-mod)/10 {
			return 0
		}
		if ans < 0 && ans < (math.MinInt32-mod)/10 {
			return 0
		}
		// if ans*10/10 != ans { // 这样也可以判断是否会溢出
		// 	return 0
		// }
		ans = ans*10 + mod
	}
	return ans
}

func myAtoi(s string) int {
	var (
		i    int
		abs  int
		sign = 1
	)
	// 去除空格
	for i < len(s) && s[i] == ' ' {
		i++
	}
	// 	判断符号
	if i < len(s) {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}

	// 	累计
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		c := int(s[i] - '0')
		abs = abs*10 + c
		// 判断是否会溢出
		if abs > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
		i++
	}
	return abs * sign
}

func isPalindrome(x int) bool {
	// 负数肯定不是回文
	if x < 0 {
		return false
	}
	// 如果以0结尾 那么除了0其他都不可能是回文
	if x%10 == 0 && x != 0 {
		return false
	}
	// 按个反转进行对比
	var r int
	for x > r {
		mod := x % 10
		x /= 10
		r = r*10 + mod
	}
	// 如果长度为奇数那么势必有个会大10倍 这里两种情况都算ok
	// 比如 12321 反转结果肯定是12和123 实际上是回文
	return x == r || x == r/10
}

// 超时
// 因为将 -2147483648 转成正数会越界，但是将 2147483647 转成负数，则不会
// 所以，我们将 a 和 b 都转成负数
// 时间复杂度：O(n)，n 是最大值 2147483647 --> 10^10
// 推荐讲解：https://leetcode-cn.com/problems/divide-two-integers/solution/jian-dan-yi-dong-javac-pythonjs-liang-sh-ptbw/
func divide(a int, b int) int {
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}

	sign := 1
	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		sign = -1
	}
	// 因为负数会比正数大一 为了避免边界情况 这里都转成负数来算
	if a > 0 {
		a = -a
	}
	if b > 0 {
		b = -b
	}

	res := 0
	for a <= b {
		var d = b
		var k = 1
		for a <= d+d {
			d += d
			k += k
		}
		a -= d
		res += k
	}
	return sign * res
}

func addBinary(a string, b string) string {
	var (
		la    = len(a) - 1
		lb    = len(b) - 1
		carry int
		res   []string
	)
	for la >= 0 || lb >= 0 {
		var ca, cb int
		if la >= 0 {
			ca = int(a[la] - '0')
			la--
		}
		if lb >= 0 {
			cb = int(b[lb] - '0')
			lb--
		}

		sum := ca + cb + carry
		carry = sum / 2
		sum %= 2
		res = append([]string{strconv.Itoa(sum)}, res...)
	}
	if carry != 0 {
		res = append([]string{strconv.Itoa(carry)}, res...)
	}
	return strings.Join(res, "")
}

func countBits(n int) []int {
	var ans = make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = oneCount(i)
	}
	return ans
}

func oneCount(n int) int {
	var ans int
	for n > 0 {
		n &= n - 1 // 将n的最后一位1变成0，比如 111 处理后就是 111&110=110 110就是110&101=100 100就是100&011=000
		ans++
	}
	return ans
}

func oneCount2(n int) int {
	var ans int
	for i := 0; i < 64; i++ {
		if n>>i&1 == 1 {
			ans++
		}
	}
	return ans
}

/*
如果输入是：nums = [2,2,3,2]，那么它的各个元素对应的32位二进制数就是
[
0000...010,
000...0010,
000..00011,
0000...010
]；
接着，对这些二进制数的对应位进行求和，得到：[00..0041]；
对这个求和结果的每一位进行3的取模运算，得到：[00..0011]；
把上面的结果从二进制转换为十进制，就是：3。这就是我们的答案。
// 题解：https://leetcode-cn.com/problems/single-number-ii/solution/ti-yi-lei-jie-wei-yun-suan-yi-wen-dai-ni-50dc/
*/
func singleNumber(nums []int) int {
	var ans int
	for i := 0; i < 31; i++ {
		var sum int
		for _, v := range nums {
			sum += (v >> i) & 1 // 提取从右往左数第i位的数值，将所有nums[i]
		}
		if sum%3 == 1 { // 如果没办法被3整除，那么说明落单的那个数的第i位是1不是0
			ans |= 1 << i
		}
	}
	return ans
}
