package goadvanced

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRand(t *testing.T) {
	p := rand.Perm(10) // 生成一个伪随机切片
	fmt.Printf("perm: %+v\n", p)
	a := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	rand.Shuffle(len(a), func(i, j int) { // 洗牌
		a[i], a[j] = a[j], a[i]
	})
	fmt.Println(a)
}
