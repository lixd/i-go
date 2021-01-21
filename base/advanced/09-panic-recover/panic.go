package main

func A() {
	defer A1()
	defer A2()
	panic("panic A")
}

func A1() {
	panic("panic A1")
}
func A2() {

}

func main() {
	A()
}
