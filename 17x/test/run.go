package main

import (
	"encoding/json"
	"fmt"
)

var MonthStr = map[int]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}

func main() {
	// var item bson.D
	// for _, v := range []int{1, 2, 3} {
	// 	item = append(item, bson.E{Key: MonthStr[v], Value: 0})
	// }
	// update := bson.D{{"$set", item}}
	// update2 := bson.D{{"$set", bson.D{
	// 	{"January", 0},
	// 	{"February", 0},
	// 	{"March", 0},
	// }}}
	// fmt.Println(update)
	// fmt.Println(update2)
	a := A{
		A1: "A1",
		A2: "A2",
		B: B{
			B1: "B1",
			B2: "B2",
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
