package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := GetUnixTime()
	fmt.Println(t1)
}

func GetUnixTime() int64 {
	local := time.FixedZone("UTC", 8*3600)
	year, month, day := time.Now().In(local).Date()
	daytime := time.Date(year, month, day, 0, 0, 0, 0, local).Unix()
	timestamp := daytime
	return timestamp
}
