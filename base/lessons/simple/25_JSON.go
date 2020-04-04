package simple

import (
	"encoding/json"
	"fmt"
)

func main() {
	u1 := User3{"illusory", 22}
	uJson, e := json.Marshal(u1)
	if e != nil {
		fmt.Println("err:", e)
	}
	fmt.Println(string(uJson))
	var u2 User3
	e = json.Unmarshal([]byte(uJson), &u2)
	if e != nil {
		fmt.Println("err:", e)
	}

	fmt.Println(u2)

}

type User3 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
