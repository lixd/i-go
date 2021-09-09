package main

import (
	"fmt"
	"log"

	"github.com/sony/sonyflake"
)

func main() {
	st := sonyflake.Settings{}
	sf := sonyflake.NewSonyflake(st)
	for i := 0; i < 10; i++ {
		id, err := sf.NextID()
		if err != nil {
			log.Println("err:", err)
		}
		fmt.Println(id)
	}
}
