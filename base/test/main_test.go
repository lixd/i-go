package main

import (
	"i-go/demo/model"
	"i-go/demo/utils"
	"testing"
)

func TestAdd(t *testing.T) {
	user := model.User{"1233333", "123456"}
	utils.TrimForStruct(&user, "Phone", "Password")
	sum := Add(1, 2)
	if sum != 3 {
		t.Error("1 and 2 result is not 3")
	}
}

func TestMultiAdd(t *testing.T) {
	var tests = []struct {
		Arg1 int
		Arg2 int
		Sum  int
	}{
		{1, 1, 2},
		{2, 3, 5},
		{-1, 2, 1},
		{0, 1, 1}}
	for _, test := range tests {
		sum := Add(test.Arg1, test.Arg2)
		if sum != test.Sum {
			t.Errorf("Add %v and %v result is not %v", test.Arg1, test.Arg2, test.Sum)
		}
	}
}
