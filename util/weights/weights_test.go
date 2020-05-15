package main

import "testing"

func BenchmarkNewWeight(b *testing.B) {
	var weights = NewWeight()
	//	添加初始化值
	weights.add("A", 10)
	weights.add("B", 20)
	weights.add("C", 30)
	weights.add("D", 30)
	weights.add("E", 10)

	allWeights := weights.getAllWeights()
	weights.AllWeight = allWeights
	for i := 0; i < b.N; i++ {
		//	进行随机
		random := weights.random()
		//	定位item
		_ = weights.locationData(random)
	}
}
