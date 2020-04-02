package simple

import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	Name  string `json:"name"` // `json:"name"` 就是 struct tag
	Age   int    `json:"age"`
	Color string `json:"color"`
}

// 指针，slice，和map的零值都是nil，即还没有分配空间，使用前需要make
type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	slice  []int
	map1   map[string]string
}

// 返回的是局部变量的地址
// 该对象是在堆还是栈上
// go语言编译器会自动决定把一个变量放在栈还是放在堆，编译器会做逃逸分析(escape analysis)
// 当发现变量的作用域没有跑出函数范围，就可以在栈上，反之则必须分配在堆。
func createCat(name, color string, age int) *Cat {
	return &Cat{name, age, color}
}

type DAO interface {
	Query() string
}
type AccountDao struct {
	Name string
}
type AccountDao2 struct {
	Name string
}

func (ad *AccountDao) Query() string {
	return ad.Name
}
func (ad2 *AccountDao2) Query() string {
	return ad2.Name
}

type AccountServer struct {
	dao *DAO
}

func main() {
	dao := AccountDao{Name: "hello"} // dao type main.AccountDao
	i := DAO(&dao)
	fmt.Printf("i type %T \n", i) // i *main.AccountDao
	as := AccountServer{&i}
	fmt.Printf("AccountServer type %T \n", as) // AccountServer type main.AccountServer
	fmt.Printf("as.dao type %T \n", as.dao)    // as.dao type *main.DAO
	fmt.Println((*as.dao).Query())

	cat1 := Cat{"喵喵", 2, "白色"}
	fmt.Println(cat1)
	var person2 Person = Person{}
	var p1 Person
	fmt.Println(p1)
	fmt.Println(person2)

	// 方式一
	var myCat1 Cat
	// 方式二
	var myCat2 Cat = Cat{"tom", 11, "灰色"}
	// 方式三
	var myCat3 *Cat = new(Cat)
	// 方式四
	var myCat4 *Cat = &Cat{}

	jsonCat, e := json.Marshal(myCat2)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("jsonCat", jsonCat)
	fmt.Println("jsonCat", string(jsonCat))
	fmt.Println(myCat1)
	fmt.Println(myCat2)
	fmt.Println(myCat3)
	fmt.Println(myCat4)

	p1.Age = 22
	p1.Name = "illusory"
	p1.Scores = [5]float64{1, 3, 4, 5, 6}
	// slice使用前 一定要make
	p1.slice = make([]int, 10)
	p1.slice[0] = 100

	// map使用前也要make
	p1.map1 = make(map[string]string)
	p1.map1["Go"] = "Golang"

	//

}
