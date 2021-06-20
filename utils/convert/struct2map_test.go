package convert

import (
	"fmt"
	"testing"
)

func TestStruct2Map(t *testing.T) {
	type user struct {
		Name    string
		Phone   string
		Address string
		Age     int
	}
	item := user{
		Name:    "China Mobile",
		Phone:   "10086",
		Address: "CQ",
		Age:     111,
	}
	m := Struct2Map(item)
	item2 := user{}
	err := Map2Struct(m, &item2)
	if err != nil {
		fmt.Println("err ", err)
	}
	fmt.Println(m)
	fmt.Printf("%#v \n", item2)
}

func BenchmarkStruct2Map(b *testing.B) {
	type user struct {
		Name    string
		Phone   string
		Address string
		Age     uint8
	}
	item := user{
		Name:    "China Mobile",
		Phone:   "10086",
		Address: "CQ",
		Age:     111,
	}
	for i := 0; i < b.N; i++ {
		_ = Struct2Map(item)
	}
}
