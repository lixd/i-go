package stack_queue_deque

// https://leetcode-cn.com/problems/valid-parentheses
// 遇到左括号就入栈 遇到右括号就和栈顶元素匹配 最后栈被清空则说明字符串有效 最后未被清空或者中途出现不匹配的情况都算无效
func isValid(s string) bool {
	var stack = make([]byte, 0)
	var m = map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			continue
		}
		// 	不是左括号则进行匹配
		value, ok := m[s[i]] // 取出当前 右括号对应的左括号
		if ok {
			// 如果栈已经空了 当前还遇到的是右括号说明也是无效的(已经找不到左括号与之匹配了)
			// 如果和栈顶元素不匹配 则字符串无效
			if len(stack) == 0 || value != stack[len(stack)-1] {
				return false
			}
			// 否则移除当前栈顶元素 继续下一轮
			stack = stack[:len(stack)-1]
		}
	}
	// 如果最后栈被清空则说明字符串有效
	return len(stack) == 0
}
func isValid2(s string) bool {
	var stack = make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, ')')
			continue
		}
		if s[i] == '[' {
			stack = append(stack, ']')
			continue
		}
		if s[i] == '{' {
			stack = append(stack, '}')
			continue
		}
		// 如果栈已经空了 当前还遇到的是右括号说明也是无效的(已经找不到左括号与之匹配了)
		// 如果和栈顶元素不匹配 则字符串无效
		if len(stack) == 0 || s[i] != stack[len(stack)-1] {
			return false
		}
		// 否则移除当前栈顶元素 继续下一轮
		stack = stack[:len(stack)-1]
	}
	// 如果最后栈被清空则说明字符串有效
	return len(stack) == 0
}
