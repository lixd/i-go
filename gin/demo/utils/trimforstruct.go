package utils

import (
	"fmt"
	"reflect"
)

func TrimForStruct(structs interface{}, string2 ...string) {
	typeOf := reflect.TypeOf(structs).Elem()
	valueOf := reflect.ValueOf(structs).Elem()
	for value := range string2 {
		for i := 0; i < typeOf.NumField(); i++ {
			if typeOf.Field(i).Name == string(value) {
				fmt.Printf("%s -- %v \n", typeOf.Field(i).Name, valueOf.Field(i).Interface())
			}
		}
	}

}
