package converse

import (
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/mitchellh/mapstructure"
	"reflect"
)

// Struct2Map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	// 如果是指针，则获取其所指向的元素
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	var data = make(map[string]interface{})
	// 只有结构体可以获取其字段信息
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			data[t.Field(i).Name] = v.Field(i).Interface()
		}
	}
	return data
}

// Map2Struct  m map s 结构体指针
func Map2Struct(m map[string]interface{}, s interface{}) error {
	decode := mapstructure.Decode(m, s)
	return decode
}

func Map2StructUtil() {
	type Ids struct {
		Id  int `json:"id"`
		Uid int `json:"uid"`
	}
	type Base struct {
		Ids
		CreateTime string `json:"create_time"`
	}
	type User struct {
		Base
		Passport string `json:"passport"`
		Password string `json:"password"`
		Nickname string `json:"nickname"`
	}
	//u:=User{
	//	Base:     Base{
	//		Ids:        Ids{
	//			Id:  1,
	//			Uid: 2,
	//		},
	//		CreateTime: "11111",
	//	},
	//	Passport: "Passport",
	//	Password: "Password",
	//	Nickname: "Nickname",
	//}
	var u User
	m := map[string]interface{}{
		"id":          "1",
		"uid":         "2",
		"create_time": "11111111",
		"passport":    "passport",
		"password":    "password",
		"nickname":    "nickname",
	}
	gconv.StructDeep(m, &u)
	fmt.Printf("%#v \n", u)
}
