package main

import (
	"debug/elf"
	"log"
)

// var d uint8
//
// func main() {
// 	var a uint8 = 1
// 	a = 2
// 	if true {
// 		a = 3
// 	}
// 	d = a
// }

// 查看SSA初始及其后续优化阶段生成的代码片段：GOSSAFUNC=main GOOS=linux GOARCH=amd64 go tool compile  main.go
// func main() {
// 	fmt.Println("123")
// }

// 查看 elf 文件信息
func main() {
	f, err := elf.Open("main")
	if err != nil {
		return
	}
	for _, section := range f.Sections {
		log.Println(section)
	}
}
