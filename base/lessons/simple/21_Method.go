package simple

import "fmt"

func main() {
	p1 := Personm{"illusory", 11}
	p1.run() //调用方法
	//实现了String()方法 fmt.Println时会自动调用String()方法输出 类似Java中的toString()
	fmt.Println(p1)
	op := OtherPeople{People{12, "male", "t"}}
	op.PeInfo()
}

type Personm struct {
	Name string
	Age  int
}

//给Personm绑定一个方法
func (p Personm) run() {
	fmt.Printf("%s running....\n", p.Name)
}

func (personm *Personm) String() string {
	str := fmt.Sprintf("Name=%v Age=%v", personm.Name, personm.Age)
	return str
}

type People struct {
	Age    int
	gender string
	Name   string
}

type OtherPeople struct {
	People
}

func (p *People) PeInfo() {
	fmt.Println("People ", p.Name, ": ", p.Age, "岁, 性别:", p.gender)
}
