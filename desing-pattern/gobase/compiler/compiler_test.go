package compiler

import (
	"fmt"
	"go/scanner"
	"go/token"
	"testing"
)

// 词法解析：将源文件进行 token 化
func TestScanner(t *testing.T) {
	// 模拟解析源文件符号化
	src := []byte("cos(x) + si*sin(x) // Euler")
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, scanner.ScanComments)
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}
