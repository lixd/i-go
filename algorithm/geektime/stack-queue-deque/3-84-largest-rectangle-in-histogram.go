package stack_queue_deque

// https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
func largestRectangleArea(heights []int) int {
	if len(heights) <= 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}
	var (
		stack         = make([]flag, 0, len(heights))
		maxArea, area int
	)
	stack = append(stack, flag{index: -1, value: -1})
	for i, v := range heights {
		if v >= stack[len(stack)-1].value {
			stack = append(stack, flag{
				index: i,
				value: v,
			})
		} else {
			for j := len(stack) - 1; j > 1; j-- {
				if stack[j].value <= v {
					break
				}
				left := stack[j-1]
				right := stack[j]
				area = (right.index - left.index + 1) * left.value
				maxArea = max(area, maxArea)
				stack = stack[:len(stack)-1]
			}
		}
	}
	for stack[len(stack)-1].value != -1 {
		right := stack[len(stack)-1]
		var index flag
		for k := len(stack) - 1; k > 0; k-- {
			if stack[k].value < right.value {
				index = stack[k]
			}
		}
		left := stack[len(stack)-2]
		area = (right.index - index.index + 1) * left.value
		maxArea = max(area, maxArea)
		stack = stack[:len(stack)-1]
	}
	return maxArea
}

type flag struct {
	index int
	value int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea2(heights []int) int {

	res := 0

	var q []int // 通过slice维护一个单调递增栈

	for i := 0; i < len(heights); i++ {
		for len(q) > 0 && heights[i] < heights[q[len(q)-1]] {
			h := heights[q[len(q)-1]] // 以出栈元素为高，计算最大矩形的面积
			q = q[:len(q)-1]

			var w int // 计算宽
			if len(q) == 0 {
				w = i
			} else {
				w = i - q[len(q)-1] - 1
			}

			s := h * w
			if s > res {
				res = s
			}
		}

		q = append(q, i)
	}

	// 清空栈内元素,确保以每个元素作为高，并计算其面积
	for len(q) > 0 {
		h := heights[q[len(q)-1]] // 以出栈元素为高，计算最大矩形的面积
		q = q[:len(q)-1]

		var w int // 计算宽
		if len(q) > 0 {
			w = len(heights) - q[len(q)-1] - 1
		} else {
			w = len(heights)
		}

		s := h * w
		if s > res {
			res = s
		}
	}

	return res
}
