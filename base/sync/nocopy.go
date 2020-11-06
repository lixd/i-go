package main

import "fmt"

// 需先定义noCopySign的结构体,该结构体名字随意,关键是要实现 sync.Locker 接口
type noCopySign struct{}

func (*noCopySign) Lock()   {}
func (*noCopySign) Unlock() {}

// 这里定义自己的结构体,包含上面的noCopySign结构体即可
type nocopyTest struct {
	nocopy noCopySign
	a      int32
}

func main() {
	///		错误做法	///
	t1 := nocopyTest{}
	t1.a = 111
	fmt.Printf("t1:%v \n", t1)

	t2 := t1
	t2.a = 222
	fmt.Printf("t1:%v,t2:%v \n", t1, t2)

	///		正确做法	///
	t3 := &nocopyTest{} // 一定要使用指针
	t3.a = 333
	fmt.Printf("t3:%v \n", t3)

	t4 := t3 // t4指向t3所指向的对象,所以是同一个对象
	t4.a = 444
	fmt.Printf("t3:%v,t4:%v \n", t3, t4)
}
