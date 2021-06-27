package main

import "time"

func main() {
	var (
		A, B int
	)
	go func() {
		A = 1
		print(B)
	}()

	go func() {
		B = 1
		print(A)
	}()
	time.Sleep(time.Second)
}
