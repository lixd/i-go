package common

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack(10)

	// for i := int64(0); i < 11; i++ { // stack is full
	for i := int64(0); i < 10; i++ {
		err := s.Push(i)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("push:", i)
	}

	// for i := 0; i < 11; i++ { // stack is empty
	for i := 0; i < 10; i++ {
		pop, err := s.Pop()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("pop:", pop)
	}
}
