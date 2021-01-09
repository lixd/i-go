package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// Str1()

	str := "hello"
	buf := String2Bytes(str)
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&str))
	fmt.Printf("%v %#v \n", &buf[0], sh.Data) // 相同的地址

	// buf := []byte{'h', 'e', 'l', 'l', 'o'}
	// str := Bytes2String(buf)
	// sh := (*reflect.SliceHeader)(unsafe.Pointer(&str))
	// fmt.Printf("%v %#v \n", &buf[0], sh.Data) // 相同的地址
}

func Str1() {
	str := "hello指月"
	fmt.Printf("%c \n", str[1]) // e
	// 转为 []byte 会重新分配内存并拷贝原数据，所以可以修改其中的内容了
	bs := []byte(str)
	bs[1] = 'z'
	fmt.Printf("%c \n", bs[1]) // z
}

func String2Bytes(str string) []byte {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&str))
	// slice 比 string 多一个 cap 属性 这里给 cap 单独赋值
	sh.Cap = sh.Len
	return *(*[]byte)(unsafe.Pointer(sh))
}

func Bytes2String(buf []byte) string {
	return *(*string)(unsafe.Pointer(&buf))
}
