package test

import "time"

func addUpper(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
		time.Sleep(time.Millisecond * 50)
	}
	return sum

}
