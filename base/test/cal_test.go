package test

import (
	"testing"
)

func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		//fmt.Printf("AddUpper(10) error 期望值%v 实际值%v ",55,res)
		t.Fatalf("AddUpper(10) error 期望值%v 实际值%v ", 55, res)
	}

	//正确
	t.Log("AddUpper(10) 执行正确")
}
