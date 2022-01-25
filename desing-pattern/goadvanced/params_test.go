package goadvanced

import "testing"

func TestPrint(t *testing.T) {
	a := []interface{}{"123", "abc"}
	Print(a) // [123 abc] 等价于 Print([]interface{}{"123", "abc"})
	Print([]interface{}{"123", "abc"})
	Print(a...) // 123 abc 等价于 Print("123","abc")
	Print("123", "abc")
}
