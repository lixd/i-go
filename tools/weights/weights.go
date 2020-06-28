package main

import (
	"math/rand"
	"time"
)

// 权重随机算法 简单实现
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
func (w *Weights) GetAllWeights() int {
	return 100
}

// NewWeight
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
func (w *Weights) Add(data string, weights int) {
	item := WeightsItem{
		Data:          data,
		UseWeights:    w.UseWeights,
		CurrentWeight: weights,
	}
	w.Data = append(w.Data, item)
	w.UseWeights += weights
}

// random 根据总权重进行随机
func (w *Weights) Random() int {
	if w.AllWeight == 0 {
		return 1
	}
	// 这里需要+1 因为可能会随机到0 如果权重为0肯定是不会被选到的 需要处理一下
	return w.Rand.Intn(w.AllWeight) + 1
}

// locationData 根据随机值 定位具体的item
func (w *Weights) LocationData(rand int) *WeightsItem {
	for _, v := range w.Data {
		if rand > v.UseWeights && rand <= v.UseWeights+v.CurrentWeight {
			return &v
		}
	}
	return nil
}
