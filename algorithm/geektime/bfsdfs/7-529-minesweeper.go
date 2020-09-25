package bfsdfs

// https://leetcode-cn.com/problems/minesweeper/|
// 通过dirX dirY两个切片来控制xy变换，从而找到相邻的8个块
var dirX = []int{0, 1, 0, -1, 1, 1, -1, -1}
var dirY = []int{1, 0, -1, 0, 1, -1, 1, -1}

func updateBoard(board [][]byte, click []int) [][]byte {
	x, y := click[0], click[1]
	// 是地雷直接游戏结束
	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}
	dfsMines(board, x, y)
	return board
}

func dfsMines(board [][]byte, x, y int) {
	cnt := 0
	for i := 0; i < 8; i++ {
		tx, ty := x+dirX[i], y+dirY[i]
		if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) {
			continue
		}
		// 不用判断 M，因为如果有 M 的话游戏已经结束了
		if board[tx][ty] == 'M' {
			cnt++
		}
	}
	if cnt > 0 {
		board[x][y] = byte(cnt + '0')
	} else {
		board[x][y] = 'B'
		for i := 0; i < 8; i++ {
			tx, ty := x+dirX[i], y+dirY[i]
			// 这里不需要在存在 B 的时候继续扩展，因为 B 之前被点击的时候已经被扩展过了
			if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) || board[tx][ty] != 'E' {
				continue
			}
			dfsMines(board, tx, ty)
		}
	}
}
