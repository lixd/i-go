package main

import (
	"fmt"
	"math"
)

type ShapeInterface interface {
	Area() float64
	GetName() string
}

// 标准形状，它的面积为0.0
type Shape struct {
	name string
}

func (s *Shape) Area() float64 {
	return 0.0
}

func (s *Shape) GetName() string {
	return s.name
}

func PrintArea(s ShapeInterface) {
	fmt.Printf("%s : Area %v\r\n", s.GetName(), s.Area())
}

// 矩形 : 重新定义了Area方法
type Rectangle struct {
	Shape
	w, h float64
}

func (r *Rectangle) Area() float64 {
	return r.w * r.h
}

// 圆形  : 重新定义 Area 和PrintArea 方法
type Circle struct {
	Shape
	r float64
}

func (c *Circle) Area() float64 {
	return c.r * c.r * math.Pi
}

func (c *Circle) PrintArea() {
	fmt.Printf("%s : Area %v\r\n", c.GetName(), c.Area())
}
func main() {
	r := Rectangle{Shape: Shape{name: "Rectangle"}, w: 5, h: 4}
	PrintArea(&r)
}
