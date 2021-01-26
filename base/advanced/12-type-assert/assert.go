package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	eface2concrete()
	iface2concrete()
	eface2iface()
	iface2iface()
}

func eface2concrete() {
	var e interface{}
	f, _ := os.Open("17x.txt")
	e = f
	// e = "17x"
	r, ok := e.(*os.File)
	if ok {
		fmt.Println(r)
	}
}

func iface2concrete() {
	var rw io.ReadWriter
	f, _ := os.Open("17x.txt")
	rw = f
	// rw = &MyRW{}
	r, ok := rw.(*os.File)
	if ok {
		fmt.Println(r)
	}
}

type MyRW struct {
}

func (rw *MyRW) Read(p []byte) (n int, err error)  { return 0, err }
func (rw *MyRW) Write(p []byte) (n int, err error) { return 0, err }

func eface2iface() {
	var e interface{}
	f, _ := os.Open("17x.txt")
	e = f
	r, ok := e.(io.ReadWriter)
	if ok {
		fmt.Println(r)
	}
}
func iface2iface() {
	var w io.Writer
	f, _ := os.Open("17x.txt")
	w = f
	r, ok := w.(io.ReadWriter)
	if ok {
		fmt.Println(r)
	}
}
