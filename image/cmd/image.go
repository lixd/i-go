package main

import (
	"i-go/image/core"
	"os"
)

const (
	path = "D:\\usr\\img\\origin.png"
	save = "D:\\usr\\img\\gray.png"
)

func main() {
	//gray()
	saturation()
}
func gray() {
	// 1.读取数据
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 2.灰度化
	grays := core.RGB2Gray(file)
	// 3.保存
	create, err := os.Create(save)
	if err != nil {
		panic(err)
	}
	defer create.Close()
	create.Write(grays)
}
func saturation() {
	// 1.读取数据
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 2.灰度化
	grays := core.AdjustSaturation(file, 0.6)
	// 3.保存
	create, err := os.Create(save)
	if err != nil {
		panic(err)
	}
	defer create.Close()
	create.Write(grays)
}
