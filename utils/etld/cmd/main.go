package main

import (
	"fmt"

	"i-go/utils/etld"
)

func main() {
	parse := etld.Parse("www.lixueduan.com")
	fmt.Printf("%+v\n", parse)
}
