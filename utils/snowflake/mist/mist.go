/*
* 薄雾算法
*
* 1      2                                                     48         56       64
* +------+-----------------------------------------------------+----------+----------+
* retain | incr                                                |   salt   | sequence |
* +------+-----------------------------------------------------+----------+----------+
* 0      | 0000000000 0000000000 0000000000 0000000000 0000000 | 00000000 | 00000000 |
* +------+-----------------------------------------------------+------------+--------+
*
* 0. 最高位，占 1 位，保持为 0，使得值永远为正数；
* 1. 自增数，占 47 位，自增数在高位能保证结果值呈递增态势，遂低位可以为所欲为；
* 2. 随机因子一，占 8 位，上限数值 255，使结果值不可预测；
* 3. 随机因子二，占 8 位，上限数值 255，使结果值不可预测；
*
* 编号上限为百万亿级，上限值计算为 140737488355327 即 int64(1 << 47 - 1)，假设每天取值 10 亿，能使用 385+ 年
 */

package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
)

const saltBit = uint(8)              // 随机因子二进制位数
const saltShift = uint(8)            // 随机因子移位数
const incShift = saltBit + saltShift // 自增数移位数

type Mist struct {
	sync.Mutex       // 互斥锁
	inc        int64 // 自增数
	saltA      int64 // 随机因子一
	saltB      int64 // 随机因子二
}

// NewMist 初始化 Mist 结构体
func NewMist() *Mist {
	mist := Mist{inc: 1}
	return &mist
}

// Generate 生成唯一编号
func (c *Mist) Generate() int64 {
	c.Lock()
	c.inc++
	// 获取随机因子数值 ｜ 使用真随机函数提高性能
	randA, _ := rand.Int(rand.Reader, big.NewInt(255))
	c.saltA = randA.Int64()
	randB, _ := rand.Int(rand.Reader, big.NewInt(255))
	c.saltB = randB.Int64()
	// 通过位运算实现自动占位
	mist := c.inc<<incShift | (c.saltA << saltShift) | c.saltB
	c.Unlock()
	return mist
}

func main() {
	// 使用方法
	mist := NewMist()
	for i := 0; i < 10000; i++ {
		fmt.Println(mist.Generate())
	}
}
