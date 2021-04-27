package main

import (
	"fmt"
	"math"
)

// GO 伪多态的实现方式 将接口类型作为参数传递
func main() {
	s := shape{name: "shape"}
	printArea(&s)
	r := rectangle{shape: shape{name: "rectangle"}, w: 5, h: 4}
	printArea(&r)
	c := circle{shape: shape{name: "rectangle"}, r: 5}
	printArea(&c)
}

type iShape interface {
	Area() float64
	GetName() string
}

func printArea(s iShape) {
	fmt.Printf("%s : Area %v\r\n", s.GetName(), s.Area())
}

// shape 标准形状，它的面积默认为0.0
type shape struct {
	name string
}

func (s *shape) Area() float64 {
	return 0.0
}

func (s *shape) GetName() string {
	return s.name
}

// 矩形 : 重新定义了Area方法
type rectangle struct {
	shape
	w, h float64
}

func (r *rectangle) Area() float64 {
	return r.w * r.h
}

// circle 圆 πr^2
type circle struct {
	shape
	r float64
}

func (c *circle) Area() float64 {
	return c.r * c.r * math.Pi
}
