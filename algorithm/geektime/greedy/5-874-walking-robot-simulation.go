package greedy

import (
	"fmt"
	"math"
)

// https://leetcode-cn.com/problems/walking-robot-simulation/description/
func robotSim(commands []int, obstacles [][]int) int {
	// 方向：上右下左0123
	var (
		result     = 0                     // 结果
		curDir     = 0                     // 当前方向
		mObstacles = make(map[string]bool) //障碍物哈希
	)
	x, y, stepX, stepY := 0, 0, []int{0, 1, 0, -1}, []int{1, 0, -1, 0} // 当前的位置 以及 xy轴上各个方向移动的大小
	for _, v := range obstacles {
		mObstacles[fmt.Sprintf("%d-%d", v[0], v[1])] = true
	}
	for _, v := range commands {
		switch v {
		case -1:
			curDir = (curDir + 1) % 4
		case -2:
			curDir = (curDir + 3) % 4
		default:
			for i := 0; i < v; i++ {
				tempX, tempY := x+stepX[curDir], y+stepY[curDir]
				// 碰到障碍物，就不要移动了
				if mObstacles[fmt.Sprintf("%d-%d", tempX, tempY)] {
					break
				}
				x, y = tempX, tempY
				result = int(math.Max(float64(x*x+y*y), float64(result)))
			}
		}
	}
	return result
}
