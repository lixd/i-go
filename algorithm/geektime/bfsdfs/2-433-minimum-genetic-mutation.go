package bfsdfs

func minMutation(start string, end string, bank []string) int {
	var bankMap = map[string]bool{}
	var changes = []string{"A", "C", "G", "T"}

	for _, value := range bank {
		bankMap[value] = true
	}
	if len(start) != len(end) || !isInBank(end, bankMap) {
		return -1
	}
	// 采用广度优先的方式
	return bfs(0, start, end, bankMap, changes)
}
func bfs(changeTimes int, start string, end string, bankMap map[string]bool, changes []string) int {
	if start == end {
		return changeTimes
	}
	var curLevelQueue []string
	var nextLevelQueue []string
	var visited = make(map[string]bool)
	curLevelQueue = append(curLevelQueue, start)
	for len(curLevelQueue) != 0 {
		node := curLevelQueue[0]
		visited[node] = true
		for i := 0; i < len(node); i++ {
			for _, change := range changes {
				newNode := node[:i] + change + node[i+1:]
				if _, ok := visited[newNode]; !ok && isInBank(newNode, bankMap) {
					if newNode == end {
						return changeTimes + 1
					} else {
						visited[newNode] = true
						nextLevelQueue = append(nextLevelQueue, newNode)
					}
				}
			}
		}
		curLevelQueue = curLevelQueue[1:]
		if len(curLevelQueue) == 0 {
			curLevelQueue = nextLevelQueue
			nextLevelQueue = []string{}
			changeTimes++
		}
	}
	return -1
}
func isInBank(str string, bankMap map[string]bool) bool {
	if _, ok := bankMap[str]; ok {
		return true
	}
	return false
}

func minMutation2(start string, end string, bank []string) int {
	used := make([]int, len(bank))
	stack := []string{start}
	nums := 0
	for len(stack) != 0 {
		nums++
		level := len(stack)
		for i := 0; i < level; i++ {
			for k, v := range bank {
				// 遍历基因库 寻找与目标串只差一位的字符串
				if v != start && used[k] == 0 {
					diff := 0
					for m := 0; m < 8; m++ {
						if v[m] != stack[i][m] {
							diff++
						}
						if diff > 1 {
							break
						}
					}
					if diff == 1 {
						used[k] = nums
						if bank[k] == end {
							return nums
						}
						stack = append(stack, bank[k])
					}
				}
			}
		}
		stack = stack[level:]
	}
	return -1
}
