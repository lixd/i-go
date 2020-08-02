package array_linkedlist_skiplist

// https://leetcode-cn.com/problems/plus-one
/*
考虑进位，从数组最后一位开始往前遍历
如果遇到9的，那么+1之后变10，所以该位置写0，并继续往前走，
如果走到数组index=0的地方，说明走完了，则扩展数组，然后跳过本次循环
其他情况如果遇见非9的数字，然后在该位置+1
*/
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			// 该位为 9 则需要进位 先把该位 置0
			digits[i] = 0
			if i == 0 {
				// 如果已经是第一位了 说明需要扩大数组 且最前面一位为 1
				// 所以直接加到为 1 的数组上去
				digits = append([]int{1}, digits...)
			}
			continue
		} else {
			// 不为 9 则不需要进位 直接 +1 即可退出循环
			digits[i]++
			break
		}
	}

	return digits
}
