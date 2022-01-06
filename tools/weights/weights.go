package main

import (
	"math/rand"
	"time"
)

// 权重随机算法 简单实现
/*
1）根据权重将每个item对应到一个取值范围
2）然后生成一个随机数
3）该随机数在哪个范围就选选取到哪个item
*/

type Weights struct {
	list        []Item // 所有项目
	totalWeight int64  // 总权重
	r           *rand.Rand
}

type Item struct {
	Key    string  // 当前项目名
	Weight int64   // 当前项目权重
	Range  []int64 // 当前项目所占区间(左闭右开区间)
}

// NewWeight new 一个 weight 实例
func NewWeight() *Weights {
	var weights = &Weights{
		list:        nil,
		totalWeight: 0,
		r:           rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	return weights
}

// Add 添加item
func (w *Weights) Add(key string, weight int64) {
	if weight == 0 {
		return
	}
	item := Item{
		Key:    key,
		Weight: weight,
		Range:  []int64{w.totalWeight, w.totalWeight + weight},
	}
	w.list = append(w.list, item)
	w.totalWeight += weight
}

// Random 根据总权重进行随机
func (w *Weights) Random() int64 {
	if w.totalWeight == 0 {
		return 1
	}
	return w.r.Int63() % w.totalWeight
}

// LocationData 根据数值 定位具体的item
func (w *Weights) LocationData(randVal int64) Item {
	for _, v := range w.list {
		// 随机值在该item的范围内则对应该item
		if randVal >= v.Range[0] && randVal < v.Range[1] {
			return v
		}
	}
	return Item{}
}
