package lcof

// 使用 dfs + 回溯，走过的位置做标记防止重复递归
func exist(board [][]byte, word string) bool {
	// 优先处理特殊情况
	if word == "" {
		return true
	}
	if len(board) == 0 || len(board[0]) == 0 {
		return true
	}

	// 创建一个二维数组used来记录使用过的元素
	raw, col := len(board), len(board[0])
	used := make([][]bool, raw)
	for i := 0; i < raw; i++ {
		used[i] = make([]bool, col)
	}

	for i := 0; i < raw; i++ {
		for j := 0; j < col; j++ {
			// 	找到第一个满足word的元素开始dfs
			if board[i][j] == word[0] {
				// dfs 之前标记该元素已处理
				used[i][j] = true
				ok := dfs(board, used, i, j, word[1:]) // word[1:] 因为前一个已经找到了,下次递归从后一个字符开始
				if ok {
					return true
				}
				// dfs之后将标记修改为未处理
				used[i][j] = false
			}
		}
	}
	return false
}

// ds 用于控制方向 比如{-1,0}表示row+(-1),col+0 即向下走
var ds = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func dfs(board [][]byte, used [][]bool, row, col int, word string) bool {
	// word为空说明已经找完了 直接返回 true
	if word == "" {
		return true
	}
	var r, c int
	ds := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range ds {
		r, c = row+d[0], col+d[1]
		// r >= 0 && r < len(board) && c >= 0 && c < len(board[0]) 保证数组不会越界
		// !used[r][c] 保证不会重复递归
		// board[r][c] == word[0] 说明又找到了当前要匹配的元素 进入下一轮递归
		if r >= 0 && r < len(board) && c >= 0 && c < len(board[0]) && !used[r][c] && board[r][c] == word[0] {
			// 同样的处理逻辑，递归前修改used标记为true 防止重复递归
			used[r][c] = true
			if dfs(board, used, r, c, word[1:]) {
				return true
			}
			// 递归结束将标记修改回去,这样不影响下次递归
			used[r][c] = false
		}
	}
	return false
}
