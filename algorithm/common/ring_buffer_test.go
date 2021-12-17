package common

import (
	"fmt"
	"testing"
)

func TestNewRingBuffer(t *testing.T) {
	buffer := NewRingBuffer(10)
	for i := int64(0); i < 10; i++ {
		err := buffer.Write(i)
		if err != nil {
			t.Fatal(err)
		}
	}
	for i := 0; i < 10; i++ {
		read, err := buffer.Read()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("read:", read)
	}

	for i := int64(10); i < 20; i++ {
		err := buffer.Write(i)
		if err != nil {
			t.Fatal(err)
		}
	}
	for i := 0; i < 10; i++ {
		read, err := buffer.Read()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("read:", read)
	}
}
