package main

import (
	"fmt"
	"unsafe"
)

// 字节对齐
/*
1.结构体的成员变量，第一个成员变量的偏移量为 0。往后的每个成员变量的对齐值必须为编译器默认对齐长度（#pragma pack(n)）或当前成员变量类型的
长度（unsafe.Sizeof），取最小值作为当前类型的对齐值。其偏移量必须为对齐值的整数倍
2.结构体本身，对齐值必须为编译器默认对齐长度（#pragma pack(n)）或结构体的所有成员变量类型中的最大长度，取最大数的最小整数倍作为对齐值
3.结合以上两点，可得知若编译器默认对齐长度（#pragma pack(n)）超过结构体内成员变量的类型最大长度时，默认对齐长度是没有任何意义的
*/
func main() {
	printBase()

	fmt.Printf("struct size:%d align: %d\n", unsafe.Sizeof(Part1{}), unsafe.Alignof(Part1{}))
}

// 不同类型占的大小和对齐系数 最大为8应该64位机器默认对其系数位8
func printBase() {
	fmt.Printf("bool size: %d align: %d\n", unsafe.Sizeof(bool(true)), unsafe.Alignof(bool(true)))
	fmt.Printf("byte size: %d align: %d\n", unsafe.Sizeof(byte(0)), unsafe.Alignof(byte(0)))

	fmt.Printf("int8 size: %d align: %d\n", unsafe.Sizeof(int8(0)), unsafe.Alignof(int8(0)))
	fmt.Printf("int32 size: %d align: %d\n", unsafe.Sizeof(int32(0)), unsafe.Alignof(int32(0)))
	fmt.Printf("int64 size: %d align: %d\n", unsafe.Sizeof(int64(0)), unsafe.Alignof(int64(0)))
	fmt.Printf("int size: %d align: %d\n", unsafe.Sizeof(int(0)), unsafe.Alignof(int(0)))

	fmt.Printf("float32 size: %d align: %d\n", unsafe.Sizeof(float32(0)), unsafe.Alignof(float32(0)))
	fmt.Printf("float64 size: %d align: %d\n", unsafe.Sizeof(float64(0)), unsafe.Alignof(float64(0)))

	fmt.Printf("string size: %d align: %d\n", unsafe.Sizeof(string("illusory")), unsafe.Alignof(string("illusory")))
	//map和类型无关 都是8位
	fmt.Printf("string map size: %d align: %d\n", unsafe.Sizeof(map[string]string{}), unsafe.Alignof(map[string]string{}))
	fmt.Printf("int map size: %d align: %d\n", unsafe.Sizeof(map[int64]int64{}), unsafe.Alignof(map[string]string{}))
	//切片和类型无关 都是24位
	fmt.Printf("string slice size: %d align: %d\n", unsafe.Sizeof([]string{"illusory"}), unsafe.Alignof([]string{"illusory"}))
	fmt.Printf("int8 slice size: %d align: %d\n", unsafe.Sizeof([]int8{1}), unsafe.Alignof([]int8{1}))
	//array 为16
	fmt.Printf("string array size: %d align: %d\n", unsafe.Sizeof([1]string{"illusory"}), unsafe.Alignof([1]string{"illusory"}))
}

type Part1 struct {
	A bool              // 长度1 偏移0 填充0 总大小1
	B byte              // 长度1 偏移1 填充0 总大小2
	C int8              // 长度1 偏移2 填充0 总大小3
	D int16             // 长度2 偏移3 填充1 总大小6
	E int32             // 长度4 偏移5 填充2 总大小12
	F int64             // 长度8 偏移11 填充4 总大小22
	G float32           // 长度4 偏移21 填充2 总大小28
	H float64           // 长度8 偏移27 填充4 总大小40
	I string            // 长度16 偏移39 填充0 总大小56
	J map[string]string // 长度8 偏移55 填充0 总大小64
	K []string          // 长度24 偏移63 填充0 总大小88
	L [1]string         // 长度16 偏移87 填充0 总大小104
	//104 刚好是8的倍数 不需要填充
}
