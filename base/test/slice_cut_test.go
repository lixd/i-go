package main_test

import (
	"fmt"
	"testing"
)

func TestSliceCut(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	k := 2
	res := append(nums[:k], nums[k+1:]...)
	// nums[:k]=[1,2] len=2 cap=4
	// nums[k+1:]=[4] len=1 cap=1
	// res=[1,2,4] len=1 cap=4
	fmt.Printf("res =%v \n", res)
	fmt.Printf("nums =%v \n", nums)
	fmt.Printf("nums[:k] len=%v cap=%v \n", len(nums[:k]), cap(nums[:k]))
	fmt.Printf("nums[k+1:] len=%v cap=%v \n", len(nums[k+1:]), cap(nums[k+1:]))
	fmt.Printf("res len=%v cap=%v \n", len(res), cap(res))
}
