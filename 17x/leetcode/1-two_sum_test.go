package leetcode

import (
	"fmt"
	"math"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{s: "dvdf"}, want: 3},
		{name: "2", args: args{s: "ohomm"}, want: 3},
		{name: "3", args: args{s: "pwwkew"}, want: 3},
		{name: "4", args: args{s: " "}, want: 1},
		{name: "5", args: args{s: "tmmzuxt"}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring2(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindrome2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{s: "babad"}, want: "bab"},
		{name: "2", args: args{s: "a"}, want: "a"},
		{name: "3", args: args{s: "ac"}, want: "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{x: 123}, want: 321},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{s: "-91283472332"}, want: -2147483648},
		{name: "2", args: args{s: ""}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_divide(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{a: 0, b: 1}, want: 0},
		{name: "2", args: args{a: -1, b: 1}, want: -1},
		{name: "3", args: args{a: math.MaxInt32, b: 1}, want: math.MaxInt32},
		{name: "4", args: args{a: math.MinInt32, b: 1}, want: math.MinInt32},
		{name: "5", args: args{a: math.MaxInt32, b: -1}, want: math.MaxInt32 / -1},
		{name: "6", args: args{a: math.MinInt32, b: -1}, want: 2147483647},
		{name: "7", args: args{a: math.MinInt32, b: 3}, want: math.MinInt32 / 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{a: "1010", b: "1011"}, want: "10101"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_oneCount2(t *testing.T) {
	fmt.Println(oneCount2(312313))
	fmt.Println(oneCount(312313))
}
