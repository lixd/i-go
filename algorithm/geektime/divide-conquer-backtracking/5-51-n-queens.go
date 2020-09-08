package recursion

// https://leetcode-cn.com/problems/n-queens/ 八皇后问题
var solutions [][]string

func solveNQueens(n int) [][]string {
	solutions = [][]string{}
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	columns := map[int]bool{}
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}
	backtracks(queens, n, 0, columns, diagonals1, diagonals2)
	return solutions
}

func backtracks(queens []int, n, row int, columns, diagonals1, diagonals2 map[int]bool) {
	if row == n {
		board := generateBoard(queens, n)
		solutions = append(solutions, board)
		return
	}
	for i := 0; i < n; i++ {
		// 判断当前列有没有
		if columns[i] {
			continue
		}
		// 判断左右对角线有没有
		diagonal1 := row - i
		if diagonals1[diagonal1] {
			continue
		}
		diagonal2 := row + i
		if diagonals2[diagonal2] {
			continue
		}
		//
		// process 当前列、左、右对角线都没有的话说明当前位置可以放 于是就加进去
		queens[row] = i
		columns[i] = true
		diagonals1[diagonal1], diagonals2[diagonal2] = true, true
		// drill down 到下一层
		backtracks(queens, n, row+1, columns, diagonals1, diagonals2)
		// reverse 回溯回来需要清理当前层改变的状态
		queens[row] = -1
		delete(columns, i)
		delete(diagonals1, diagonal1)
		delete(diagonals2, diagonal2)
	}
}

// generateBoard 根据结果生成棋盘
func generateBoard(queens []int, n int) []string {
	var board []string
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}
