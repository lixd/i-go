package main

import (
	"fmt"
	"strconv"

	"github.com/bits-and-blooms/bloom/v3"
)

/*
布隆过滤器原理:
存入数据时将数据进行 hash,将过滤器中 hash 值对应的位置置为1（一般会使用多个hash方法以降低冲突概率）。
检测数据是否存在时使用同样的 hash 方法计算出 hash 值，判断过滤器中该位置是否为1:
 1) 如果为1说明该数据可能存在与过滤器中（因为可能出现 hash 冲突，不能100%说明一定存在）
 2) 如果不为1说明该数据肯定不存在。
所以:BloomFilter不能100%确定数据在列表中(有hash冲突)，但是可以确定100%不在。
*/
func main() {
	// 参数表示:1W 个位置,1%的误差 位置越多,误差越小则占用的内容越大
	filter := bloom.NewWithEstimates(1_0000, 0.01)
	// 添加数据
	filter.Add([]byte("Golang"))
	for i := 0; i < 1000; i++ {
		filter.Add([]byte(strconv.Itoa(i)))
	}
	// 测试是否存在
	ok1 := filter.Test([]byte("1"))
	fmt.Println("1是否存在:", ok1)
	ok1001 := filter.Test([]byte("10001"))
	fmt.Println("1001是否存在:", ok1001)
}
