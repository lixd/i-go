package daily

// https://leetcode-cn.com/problems/positions-of-large-groups/
func largeGroupPositions(s string) [][]int {
	var (
		ans  = make([][]int, 0)
		arr  = make([]int, 0)
		prev int32
	)
	for i, v := range s {
		if v == prev || i == 0 {
			arr = append(arr, i)
			prev = v
		} else {
			if len(arr) >= 3 {
				ans = append(ans, []int{arr[0], arr[len(arr)-1]})
			}
			if len(arr) != 0 {
				arr = make([]int, 0)
				arr = append(arr, i)
				prev = v
			}
		}
	}
	if len(arr) >= 3 {
		ans = append(ans, []int{arr[0], arr[len(arr)-1]})
	}
	return ans
}

func largeGroupPositions2(s string) [][]int {
	var (
		ans = make([][]int, 0)
		cnt = 1
	)
	for i := range s {
		if i == len(s)-1 || s[i] != s[i+1] {
			if cnt >= 3 {
				ans = append(ans, []int{i - cnt + 1, i})
			}
			cnt = 1
		} else {
			cnt++
		}
	}
	return ans
}
