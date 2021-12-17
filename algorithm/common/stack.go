package common

import (
	"errors"
)

type stack struct {
	data []int64 // 具体存放数据的容器
	p    int64   // 当前指针所处位置，即数组的索引
	l    int64   // 栈的容量，即数组的容量
}

func NewStack(l int64) *stack {
	return &stack{
		data: make([]int64, l),
		p:    -1, // 初始指针默认在 -1 位置
		l:    l,
	}
}

// Pop 弹出元素并移动指针
func (s *stack) Pop() (int64, error) {
	if s.p < 0 {
		return 0, errors.New("stack is empty")
	}
	data := s.data[s.p]
	s.p--
	return data, nil
}

// Push 写入元素需要先将指针移动到下一个位置,然后再写入
func (s *stack) Push(data int64) error {
	if s.p == s.l-1 {
		return errors.New("stack is full")
	}
	s.p++
	s.data[s.p] = data
	return nil
}
