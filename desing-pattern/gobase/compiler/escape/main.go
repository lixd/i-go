package main

var z *int

func escape() {
	a := 1
	z = &a
}
func escape2() int {
	a := 1
	z = &a // 虽然有引用，但是没有逃逸
	return *z
}

var o *int

func escape3() {
	l := new(int)
	*l = 42
	m := &l
	n := &m
	o = **n
}

// 逃逸分析：go tool compile -m=2 main.go
