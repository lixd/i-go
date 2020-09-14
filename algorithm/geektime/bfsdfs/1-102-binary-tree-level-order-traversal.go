package bfsdfs

import "container/list"

func levelOrder(root *TreeNode) [][]int {
	var (
		result  [][]int
		current []*TreeNode
	)
	current = append(current, root)
	for {
		if len(current) == 0 {
			break
		}
		var next []*TreeNode
		var ret []int
		for _, v := range current {
			if v == nil {
				continue
			}
			// 记录下一层的所有节点
			next = append(next, v.Left)
			next = append(next, v.Right)
			// 存储当前层的值
			ret = append(ret, v.Val)
		}
		if len(ret) != 0 {
			result = append(result, ret)
		}
		// 把下一层的节点在赋值给 current 如此循环
		current = next
	}
	return result
}

// dfs 深度优先 遍历第一层的时候把值存到数组第一个元素、第二层则存到第二个元素 通过 level 来记住当前的层级即可
func dfs(root *TreeNode, level int, res [][]int) [][]int {
	if root == nil {
		return res
	}
	if len(res) == level {
		res = append(res, []int{root.Val})
	} else {
		res[level] = append(res[level], root.Val)
	}
	res = dfs(root.Left, level+1, res)
	res = dfs(root.Right, level+1, res)
	return res
}

// bfs 广度优先
func levelOrderBfs(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	// 定义一个双向队列
	queue := list.New()
	// 头部插入根节点
	queue.PushFront(root)
	// 进行广度搜索
	for queue.Len() > 0 {
		var current []int
		listLength := queue.Len()
		for i := 0; i < listLength; i++ {
			// 每次遍历都从队列尾部移除一个
			node := queue.Remove(queue.Back()).(*TreeNode)
			// 记录当前值
			current = append(current, node.Val)
			// 并将其左右儿子节点存入队列 等待下一次遍历
			if node.Left != nil {
				// 插入头部
				queue.PushFront(node.Left)
			}
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
		}
		// 退出本次循环后 当前队列的长度已经变成了下一层级的节点个数 所以最前面的条件  queue.Len() > 0 仍然满足 会继续遍历
		// 这里则只需要把当前层结果存储起来即可
		result = append(result, current)
	}
	return result
}
