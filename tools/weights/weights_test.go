package main

import (
	"fmt"
	"testing"
)

func TestNewWeight(t *testing.T) {
	weights := NewWeight()
	//	添加初始化值
	weights.Add("A", 10)
	weights.Add("B", 20)
	weights.Add("C", 30)
	weights.Add("D", 30)
	weights.Add("E", 10)
	m := make(map[string]int)
	for i := 0; i < 10000; i++ {
		item := weights.LocationData(weights.Random())
		m[item.Key]++
	}

	for k, v := range m {
		// 结果基本符合之前设定的权重
		fmt.Println(k, "", v)
	}
}

// 69.8 ns/op
func BenchmarkNewWeight(b *testing.B) {
	weights := NewWeight()
	//	添加初始化值
	weights.Add("A", 10)
	weights.Add("B", 20)
	weights.Add("C", 30)
	weights.Add("D", 30)
	weights.Add("E", 10)

	for i := 0; i < b.N; i++ {
		//	进行随机
		random := weights.Random()
		//	定位item
		_ = weights.LocationData(random)
	}
}
