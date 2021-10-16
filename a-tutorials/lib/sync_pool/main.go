package main

import (
	"fmt"
	"sync"
)

type Gopher struct {
	Name   string
	Remark [1024]byte
}

func (s *Gopher) Reset() {
	s.Name = ""
	s.Remark = [1024]byte{}
}

var gopherPool = sync.Pool{
	New: func() interface{} {
		return new(Gopher)
	},
}

func main() {
	g := gopherPool.Get().(*Gopher)
	fmt.Println("首次从 pool 里获取：", g.Name)

	g.Name = "first"
	fmt.Printf("设置 p.Name = %s\n", g.Name)
	gopherPool.Put(g)

	fmt.Println("Pool 里已有一个对象：&{first}，调用 Get: ", gopherPool.Get().(*Gopher).Name)
	fmt.Println("Pool 没有对象了，调用 Get: ", gopherPool.Get().(*Gopher).Name)
}
