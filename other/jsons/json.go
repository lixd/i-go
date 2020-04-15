package josns

import (
	defaultJson "encoding/json"
	jsoniter "github.com/json-iterator/go"
)

/*
两个json库简单对比
*/

//BenchmarkSpecial  1942 ns/op
func Default() {
	var s1 = StructOne{
		A: "A",
		B: 2,
		C: "C",
		D: "D",
		E: "E",
		F: "F",
	}
	bytes, _ := defaultJson.Marshal(s1)
	var s2 StructOne
	_ = defaultJson.Unmarshal(bytes, &s2)
}

//BenchmarkSpecial  887 ns/op
// 用法兼容标准库 且性能有较大提升
func Special() {
	var s1 = StructOne{
		A: "A",
		B: 2,
		C: "C",
		D: "D",
		E: "E",
		F: "F",
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	bytes, _ := json.Marshal(s1)
	var s2 StructOne
	_ = json.Unmarshal(bytes, &s2)
}

type StructOne struct {
	A string
	B int
	C string
	D string
	E string
	F string
}
