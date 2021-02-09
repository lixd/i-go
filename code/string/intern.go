package main

import (
	"reflect"
	"strconv"
	"unsafe"
)

// https://ms2008.github.io/2019/08/18/golang-string-interning/
// 字符串内部化
import (
	"fmt"
)

// stringPtr returns a pointer to the string data.
func stringPtr(s string) uintptr {
	return (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
}

type stringIntern map[string]string

func (si stringIntern) Intern(s string) string {
	if interned, ok := si[s]; ok {
		return interned
	}
	si[s] = s
	return s
}

func main() {
	si := stringIntern{}
	s1 := si.Intern("12")
	s2 := si.Intern(strconv.Itoa(12))
	fmt.Println(stringPtr(s1) == stringPtr(s2)) // true
}
