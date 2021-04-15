package pointer

import "fmt"

func main() {
	a := 10
	var b *int
	b = &a
	fmt.Println(b)
	fmt.Printf("a %T %d \n", a, a)
	fmt.Printf("&a %T %d \n", &a, &a)
	fmt.Printf("b %T %v \n", b, b)
	fmt.Printf("*b %T %v \n", *b, *b)

	fmt.Println("------------------------")

	s1 := Student{"illusory", 22, 1}
	s2 := Student{"cloud", 11, 0}

	var p1 *Student = &s1
	var p2 *Student = &s2
	fmt.Printf("s1 %T %v \n", s1, s1)
	fmt.Printf("s2 %T %v \n", s2, s2)
	fmt.Printf("p1 %T %v \n", p1, p1)
	fmt.Printf("p2 %T %v \n", p2, p2)
	fmt.Printf("*p1 %T %v \n", *p1, *p1)
	fmt.Printf("*p2 %T %v \n", *p2, *p2)
	fmt.Println("------------------------")
	c := 10
	fmt.Printf("before change c %d \n", c)
	Change(&c)
	fmt.Printf("after change c %d \n", c)

	fmt.Println("------------------------")
	var ptrs [4]*string
	fmt.Printf("ptrs %T %v \n", ptrs, ptrs)
	vs := []string{"a", "b", "c", "d"}
	for i := 0; i < 4; i++ {
		ptrs[i] = &vs[i]
		fmt.Printf("%T %v \n", ptrs[i], ptrs[i])
	}
	fmt.Println("------------------------")
	f := 10
	var ptr1 *int
	ptr1 = &f
	var ptr2 **int
	ptr2 = &ptr1
	fmt.Printf("f %T %d \n", f, f)
	//ptr1 *int 0xc0000560d8
	fmt.Printf("ptr1 %T %v \n", ptr1, ptr1)
	//ptr2 **int 0xc000082020
	fmt.Printf("ptr2 %T %v \n", ptr2, ptr2)
	//*ptr2 *int 0xc0000560d8
	fmt.Printf("*ptr2 %T %v \n", *ptr2, *ptr2)
}

type Student struct {
	name string
	age  int
	sex  int8
}

func Change(p *int) {
	*p = 20
}
