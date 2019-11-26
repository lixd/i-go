package main

//计算a, b的平方和
func sum(a, b int) int {
	a2 := a * a
	b2 := b * b
	c := a2 + b2

	return c
}

func main() {
	sum(1, 2)
}
