package pool

import "sync"

// List 队列
type List interface {
	All() []interface{}
	Len() int64
	LPush(conn []interface{})
	RPop() interface{}
	RPopLPush() interface{}
}

// MyList list 实现
type MyList struct {
	data []interface{}
	mu   sync.Mutex
}

// NewList 根据推送任务元数据新建一个队列
func NewList(cap int64) List {
	return &MyList{
		data: make([]interface{}, 0, cap),
	}
}

// All 返回所有站点数据
func (l *MyList) All() []interface{} {
	return l.data
}

// Len 队列长度
func (l *MyList) Len() int64 {
	return int64(len(l.data))
}

// LPush 向队尾添加元素
func (l *MyList) LPush(si []interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.data = append(l.data, si...)
}

// RPop 从队头弹出一个元素
func (l *MyList) RPop() interface{} {
	l.mu.Lock()
	defer l.mu.Unlock()
	if len(l.data) == 0 {
		return nil
	}
	info := l.data[0]
	l.data = l.data[1:]
	return info
}

// RPopLPush 从队头弹出一个元素，并重新添加会队尾
func (l *MyList) RPopLPush() interface{} {
	c := l.RPop()
	l.LPush([]interface{}{c})
	return c
}
