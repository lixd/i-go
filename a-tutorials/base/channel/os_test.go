package channel

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("before")
	m.Run()
	fmt.Println("after")
}

func TestA(t *testing.T) {
	fmt.Println("A")
}

func TestB(t *testing.T) {
	fmt.Println("B")
}
