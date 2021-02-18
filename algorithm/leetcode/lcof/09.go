package lcof

// https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
/*
栈只能先进后出，模拟成队列后无法移除尾部元素，所以需要第二个栈来辅助。
新增元素时写入栈in,然后将in栈元素出栈并压如out,这样就实现了元素逆序。然后从out弹出的第一个元素就是队尾了。

入队时写入in栈，出队时从out栈出，out栈没有就把in栈的搬过去，两个栈都没有说明队列为空。
*/

type CQueue struct {
	in  stack
	out stack
}

type stack []int

func (s *stack) Push(value int) {
	*s = append(*s, value)
}

func (s *stack) Pop() int {
	n := len(*s)
	res := (*s)[n-1]
	*s = (*s)[:n-1]
	return res
}

func Constructor() CQueue {
	return CQueue{
		in:  make([]int, 0),
		out: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.in.Push(value)
}

func (this *CQueue) DeleteHead() int {
	// out栈有值就直接弹出
	if len(this.out) > 0 {
		return this.out.Pop()
	}
	if len(this.in) > 0 {
		// 否则就把 in栈的元素搬过来
		for len(this.in) != 0 {
			this.out.Push(this.in.Pop())
		}
		return this.out.Pop()
	}
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
