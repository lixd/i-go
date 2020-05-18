package main

import "testing"

func BenchmarkNewWeight(b *testing.B) {
	var weights = NewWeight()
	//	添加初始化值
	weights.Add("A", 10)
	weights.Add("B", 20)
	weights.Add("C", 30)
	weights.Add("D", 30)
	weights.Add("E", 10)

	allWeights := weights.GetAllWeights()
	weights.AllWeight = allWeights
	for i := 0; i < b.N; i++ {
		//	进行随机
		random := weights.Random()
		//	定位item
		_ = weights.LocationData(random)
	}
}
