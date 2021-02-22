package lcof

// https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
/*
思路很简单:

越界的点不可达，返回0
访问过的点不可达，返回0
数位和大于k的点不可达，返回0
其余可达的点返回1，最后累加起来
*/
func movingCount(m int, n int, k int) int {
	// 使用二维数组记录是否访问过该位置
	vis := make([][]bool, m+1)
	for i := 0; i < len(vis); i++ {
		vis[i] = make([]bool, n+1)
	}
	return bfs(0, 0, m, n, k, vis)
}

// turn 用于控制方向 比如{-1,0}表示x+(-1),y+0 即向下走
var turn = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func bfs(x, y, m, n, k int, vis [][]bool) int {
	var sum int
	// 越界的点不可达，返回0
	if x >= m || y >= n || x < 0 || y < 0 {
		return 0
	}
	// 访问过的点不可达，返回0
	if vis[x][y] {
		return 0
	}
	// 数位和大于k的点不可达，返回0
	if sumPos(x)+sumPos(y) > k {
		return 0
	}
	// 修改当前位置标记并且总点数+1
	vis[x][y] = true
	sum += 1
	// 依次向上下左右4个方向扩展
	for _, d := range turn {
		sum += bfs(x+d[0], y+d[1], m, n, k, vis)
	}
	return sum
}

func sumPos(n int) int {
	var sum int
	for n != 0 {
		sum += n % 10
		n = n / 10
	}
	return sum
}
