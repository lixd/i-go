package bfsdfs

// https://leetcode-cn.com/problems/number-of-islands/
/*
floodfill算法
1.遍历整个二维数组 遇到1时就将其相连的1全部改为0 同时记录操作次数
2.当整个二维数组都为0时结束遍历
3.操作次数就是岛屿数
*/
func numIslands(grid [][]byte) int {
	var step int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				step++
				dfsNum(i, j, grid)
			}
		}
	}
	return step
}

func dfsNum(x, y int, grid [][]byte) {
	// terminator
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == '0' {
		return
	}
	// process
	grid[x][y] = '0'
	// drill down
	dfsNum(x-1, y, grid)
	dfsNum(x+1, y, grid)
	dfsNum(x, y-1, grid)
	dfsNum(x, y+1, grid)
}
