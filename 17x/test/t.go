package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	score := 70.007
	score -= 100
	score += 0.001
	fmt.Println("score:", score)
	fmt.Println("Floor:", math.Floor(score))
	fmt.Println("add:", score+math.Abs(math.Floor(score)))
	fmt.Println("add:", Decimal3(score+math.Abs(math.Floor(score))))
}

func Decimal3(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", value), 64)
	return value
}
