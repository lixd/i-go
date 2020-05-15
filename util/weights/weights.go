package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	//初始化全局Seed
	var count = 10000000
	var staticMap = make(map[string]int)
	// 初始化xxx对象
	var weights = NewWeight()
	allWeights := weights.getAllWeights()
	weights.AllWeight = allWeights

	//	添加初始化值
	weights.add("A", 10)
	weights.add("B", 20)
	weights.add("C", 30)
	weights.add("D", 30)
	weights.add("E", 10)

	for i := 0; i < count; i++ {
		//	进行随机
		random := weights.random()
		//	定位item
		data := weights.locationData(random)
		staticMap[data.Data] += 1
	}
	fmt.Printf("统计结果%#v \n", staticMap)
	fmt.Printf("花费时间%v \n", time.Now().Sub(start))
}

type Weights struct {
	Data       []WeightsItem // 数据
	UseWeights int           // 已分配权重
	AllWeight  int           // 总权重
	Rand       *rand.Rand
}
type WeightsItem struct {
	Data          string // 数据
	UseWeights    int    // 之前item已分配权重
	CurrentWeight int    // 当前记录权重
}

// getAllWeights 计算总权重 当前默认为100
func (w *Weights) getAllWeights() int {
	return 100
}

// getAllWeights 计算总权重 当前默认为100
func NewWeight() *Weights {
	var weights = &Weights{
		Data:       nil,
		UseWeights: 0,
		AllWeight:  0,
		Rand:       rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	return weights
}

// add 添加item
func (w *Weights) add(data string, weights int) {
	item := WeightsItem{
		Data:          data,
		UseWeights:    w.UseWeights,
		CurrentWeight: weights,
	}
	w.Data = append(w.Data, item)
	w.UseWeights += weights
}

// random 根据总权重进行随机
func (w *Weights) random() int {
	// 这里需要+1 因为可能会随机到0 如果权重为0肯定是不会被选到的 需要处理一下
	return w.Rand.Intn(w.AllWeight) + 1
}

// locationData 根据随机值 定位具体的item
func (w *Weights) locationData(rand int) *WeightsItem {
	for _, v := range w.Data {
		if rand > v.UseWeights && rand <= v.UseWeights+v.CurrentWeight {
			return &v
		}
	}
	return nil
}
