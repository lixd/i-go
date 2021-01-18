package main

import "fmt"

func main() {
	map1()
}

func map1() {
	// 字面量初始化
	h1 := map[string]string{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
	}
	fmt.Println(h1)
	// 	运行时初始化
	h2 := make(map[string]string)
	h2["k1"] = "v1"
	h2["k2"] = "v2"
	h2["k3"] = "v3"
}

type B struct {
	ID string
}
