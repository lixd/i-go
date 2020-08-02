package stack_queue_deque

// https://leetcode-cn.com/problems/sliding-window-maximum/
// window 中维护元素下标 每次加入新元素时把比新元素小的全部移除去 保证 window 中第一个元素一直是当前的最大值
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) <= 0 {
		return nil
	}
	var (
		windows = make([]int, 0, k)
		res     []int
	)
	// 遍历数组
	for i, v := range nums {
		// i-k = 当窗口满时的最小下标
		// i-k >= 窗口最左边的数，则代表窗口已经满了
		if i >= k && windows[0] <= i-k {
			windows = windows[1:]
		}

		for {
			// 如果窗口最后一个小于v，则去掉。为了维护了每次窗口最左边值最大
			if len(windows) > 0 && nums[windows[len(windows)-1]] <= v {
				windows = windows[:len(windows)-1]
			} else {
				break
			}
		}

		windows = append(windows, i)
		if i >= k-1 {
			// i >= k-1 则每次第一个元素就是最大值
			res = append(res, nums[windows[0]])
		}
	}

	return res
}

func maxSlidingWindow2(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	var (
		windows = make([]int, 0, k)
		res     = make([]int, 0, len(nums)-k)
	)
	for i, v := range nums {
		// 存满了则移除第一个数 Window 中存的是下标
		if i >= k && windows[0] <= i-k {
			windows = windows[1:]
		}
		// 	移除比当前值小的所有值 保证
		for {
			if len(windows) > 0 && nums[windows[len(windows)-1]] < v {
				windows = windows[:len(windows)-1]
			} else {
				break
			}
		}
		// 每次把当前下标存入 Windows
		windows = append(windows, i)
		if i >= k-1 {
			res = append(res, nums[windows[0]])
		}
	}
	return res
}
