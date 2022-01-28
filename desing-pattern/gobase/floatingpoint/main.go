package main

import (
	"fmt"
	"math"
)

func main() {
	var number float32 = 0.085
	fmt.Printf("Starting Number: %f\n\n", number)
	bits := math.Float32bits(number)
	binary := fmt.Sprintf("%.32b", bits)
	fmt.Printf(binary)
}
