package stack_queue_deque

import (
	"fmt"
	"testing"
)

func TestMinStack_Pop(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Pop()
	param_3 := obj.Top()
	fmt.Println(param_3)
	param_4 := obj.GetMin()
	fmt.Println(param_4)
}
