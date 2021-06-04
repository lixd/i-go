package main

import (
	"fmt"

	"github.com/willf/bloom"
)

/*
布隆过滤器原理:
存入数据时将数据进行 hash,将过滤器中 hash 值对应的位置置为1（一般会使用多个hash方法以降低冲突概率）。
检测数据是否存在时使用同样的 hash 方法计算出hash值，判断过滤器中该位置是否为1:
 1) 如果为1说明该数据可能存在与过滤器中（因为可能出现hash冲突，不能100%说明一定存在）
 2) 如果不为1说明该数据肯定不存在。
*/
func main() {
	// new一个实例 2_0000 个位置(bit) 5个hash方法
	filter := bloom.New(2_0000, 5)
	filter.Add([]byte("Golang"))          // 添加数据
	test := filter.Test([]byte("Golang")) // 测试是否存在
	fmt.Println("是否存在:", test)
}
