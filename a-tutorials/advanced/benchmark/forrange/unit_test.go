package forrange

import (
	"fmt"
	"testing"
)

func TestUnit(t *testing.T) {
	in := []interface{}{"A"}
	inte := interface{}(in)
	aa := inte.([]interface{})
	fmt.Println(aa)
	for _, v := range in {
		tmp, ok := v.(string)
		if !ok {
			fmt.Println("断言失败")
		}
		fmt.Println(tmp)
	}
}
