package utils

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSubset(t *testing.T) {
	backends := make([]string, 0, 100)
	for i := 0; i < 100; i++ {
		backends = append(backends, strconv.Itoa(i))
	}
	subset := Subset(backends, "client011", 10)
	fmt.Printf("%#v\n", subset)
}

func Test_stringHelper_GetUUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		uuid := StringHelper.GetUUID()
		fmt.Println(uuid)
	}
}
