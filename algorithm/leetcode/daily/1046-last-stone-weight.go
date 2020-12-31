package daily

import (
	"container/heap"
	"sort"
)

// https://leetcode-cn.com/problems/last-stone-weight/ O(n^2logn)
func lastStoneWeight(stones []int) int {
	for len(stones) >= 2 {
		// 快排 O(nlogn)
		sort.Sort(sort.Reverse(sort.IntSlice(stones)))
		y := stones[0]
		x := stones[1]
		// 相同则直接消掉两个
		stones = stones[2:]
		if x != y {
			// 否则把第二个剩余重量加回去
			stones = append(stones, y-x)
		}
	}
	if len(stones) != 0 {
		return stones[0]
	}
	return 0
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func (h *hp) push(v int) { heap.Push(h, v) }
func (h *hp) pop() int   { return heap.Pop(h).(int) }

// lastStoneWeight2 O(nlogn) 堆排序
func lastStoneWeight2(stones []int) int {
	q := &hp{stones}
	heap.Init(q)
	for q.Len() > 1 {
		x, y := q.pop(), q.pop()
		if x > y {
			q.push(x - y)
		}
	}
	if q.Len() > 0 {
		return q.IntSlice[0]
	}
	return 0
}
