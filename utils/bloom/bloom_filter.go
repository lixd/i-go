package main

import (
	"github.com/sirupsen/logrus"
	"github.com/willf/bloom"
)

func main() {
	// new一个实例1000*20个数据 5个hash方法
	n := uint(1000)
	filter := bloom.New(20*n, 5)
	// 添加数据
	filter.Add([]byte("Love"))
	// 测试是否存在
	test := filter.Test([]byte("Love"))
	logrus.Infof("res:%v", test)
}
