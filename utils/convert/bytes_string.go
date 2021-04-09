package convert

// Bytes和String高效转换
import (
	"reflect"
	"unsafe"
)

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func String2Bytes(str string) []byte {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&str))
	// slice 比 string 多一个 cap 属性 这里给 cap 单独赋值
	sh.Cap = sh.Len
	return *(*[]byte)(unsafe.Pointer(sh))
}
