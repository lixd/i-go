package daily

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func Test_fairCandySwap(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "1", args: struct {
			A []int
			B []int
		}{A: []int{1, 1}, B: []int{2, 2}}, want: []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fairCandySwap(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fairCandySwap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAB(t *testing.T) {
	price, err := strconv.ParseFloat("0.1", 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(price)
}
