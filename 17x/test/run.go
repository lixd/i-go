package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := A{
		A1: "A1",
		A2: "A2",
		B: B{
			B1: "B1",
			B2: "B2",
			C: C{
				C1: "C1",
				A1: "A11", // 内层字段json序列化后被覆盖 json序列化时应该是把所有字段字段展开
			},
		},
	}
	bytes, _ := json.Marshal(a)
	fmt.Println(string(bytes))
}

type A struct {
	A1 string
	A2 string
	B
}
type B struct {
	B1 string
	B2 string
	C
}
type C struct {
	C1 string
	A1 string
}
