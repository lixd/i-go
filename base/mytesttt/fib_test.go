package mytesttt

import (
	"fmt"
	"testing"
)

//func TestFib(t *testing.T) {
//	type args struct {
//		n int
//	}
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{
//		{"0", args{0}, 0},
//		{"1", args{1}, 1},
//		{"2", args{2}, 1},
//		{"3", args{3}, 2},
//		{"4", args{4}, 3},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := Fib(tt.args.n); got != tt.want {
//				t.Errorf("Fib() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }
func BenchmarkFib40Parallel(b *testing.B) {
	//RunParallel会创建出多个goroutine，并将b.N分配给这些goroutine执行， 其中goroutine数量的默认值为GOMAXPROCS。
	// 用户如果想要增加非CPU受限（non-CPU-bound）基准测试的并行性， 那么可以在RunParallel之前调用SetParallelism
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			benchmarkFib(b, 40)
		}
	})
}

func ExampleFib() {
	fmt.Println(Fib(1))
	//	Output:1
}
