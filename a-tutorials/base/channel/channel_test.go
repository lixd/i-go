package channel

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
	"unsafe"
)

func TestChannel(t *testing.T) {
	var waitGroup sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	go select1(ctx, &waitGroup)
	waitGroup.Add(1)
	time.Sleep(time.Millisecond * 500)
	cancel()
	waitGroup.Wait()
}
func select1(ctx context.Context, group *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("return~")
			group.Done()
			return
		default:
			fmt.Println("default~")
		}
		time.Sleep(time.Millisecond * 100)
	}
}

type Programmer struct {
	name     string
	language string
}

func TestPointer(t *testing.T) {
	p := Programmer{"stefno", "go"}
	fmt.Println(p)

	name := (*string)(unsafe.Pointer(&p))
	*name = "qcrao"

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language)))
	*lang = "Golang"

	fmt.Println(p)

	s := []int{5}
	s = append(s, 7)
	s = append(s, 9)
	x := append(s, 11)
	y := append(s, 12)
	fmt.Println(s, x, y)
}
