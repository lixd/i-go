package goadvanced

import (
	"log"
	"testing"
)

func TestRecoverOK(t *testing.T) {
	defer func() {
		// recover 和 panic 所在栈帧刚好差了两层，为祖父级，可以捕获到
		if r := recover(); r != nil {
			log.Println("recover:", r)
		}
	}()
	panic("aaa")
}

func TestRecoverBad1(t *testing.T) {
	// 直接调用 recover 无法捕获到
	if r := recover(); r != nil {
		log.Println("recover:", r)
	}
	panic("aaa")
}

func TestRecoverBad3(t *testing.T) {
	defer func() {
		// recover 和 panic 所在栈帧差了三层，无法捕获到
		func() {
			if r := recover(); r != nil {
				log.Println("recover:", r)
			}
		}()
	}()
	panic("aaa")
}

func TestRecoverBad2(t *testing.T) {
	// recover 和 panic 所在栈帧只差了一层，同样无法捕获到
	defer recover()
	panic("aaa")
}

func TestRecoverCustomBad(t *testing.T) {
	defer func() {
		// 由于被包装了一层，因此 recover 和 panic 所在栈帧差了三层，无法捕获到
		MyRecover()
	}()
	panic("aaa")
}

func TestRecoverCustomGood(t *testing.T) {
	// 被包装了一层，调用时也省略了 defer func 这一层，因此 recover 和 panic 所在栈帧差了两层，可以捕获到
	defer MyRecover()
	panic("aaa")
}

func MyRecover() {
	log.Println("MyRecover")
	if r := recover(); r != nil {
		log.Println("recover:", r)
	}
}
