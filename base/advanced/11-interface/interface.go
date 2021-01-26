package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	eface()
}

func eface() {
	var e interface{}
	f, _ := os.Open("17x.txt")
	e = f
	fmt.Printf("%#v \n", e)
}

func iface() {
	var rw io.ReadWriter
	f, _ := os.Open("17x.txt")
	rw = f
	fmt.Printf("%#v \n", rw)
}
