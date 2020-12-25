package daily

import (
	"fmt"
	"sync"
	"testing"
)

func Test_candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{[]int{1, 0, 2}}, want: 5},
		{name: "2", args: args{[]int{1, 2, 2}}, want: 4},
		{name: "3", args: args{[]int{1, 3, 2, 2, 1}}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy(tt.args.ratings); got != tt.want {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	var wg sync.WaitGroup
	m := make(map[int]int)
	wg.Add(2)
	go write(m, &wg)
	go write(m, &wg)
	wg.Wait()
}

func write(m map[int]int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		m[i] = i
	}
}

func TestT(t *testing.T) {
	fmt.Println(1<<3 | 1)
}
