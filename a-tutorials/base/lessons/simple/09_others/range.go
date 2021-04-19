package main

import "fmt"

func main() {
	s := []interface{}{1, 2, 3, 4, 5, 6, 7, 8}
HERE:
	fmt.Printf("arr:%v \n", s)
	for k, v := range s {
		if v == 4 || v == 6 || v == 7 {
			// temp := s[k+1:]
			fmt.Printf("k:%v v:%v \n", k, v)
			temp := make([]interface{}, 0)
			for _, v := range s[k+1:] {
				temp = append(temp, v)
			}
			fmt.Println("temp1", temp)
			s = append(s[:k], "(")
			fmt.Println(s)
			s = append(s, []interface{}{fmt.Sprintf("%d", v), fmt.Sprintf("%d", v)}...)
			fmt.Println(s)
			s = append(s, ")")
			fmt.Println(s)
			fmt.Println("temp2", temp)
			s = append(s, temp...)
			fmt.Println(s)
			goto HERE
		}
	}
}
