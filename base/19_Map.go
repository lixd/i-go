package main

import (
	"fmt"
	"sort"
)

func main() {
	//map声明1
	var names map[int]string
	//使用前make 为map分配内存空间
	names = make(map[int]string)

	names[0] = "Go"
	names[1] = "C++"
	names[1] = "C" //覆盖前面的
	names[2] = "Python"
	names[3] = "Java"
	//map[0:Go 1:C++ 2:Python 3:Java]
	fmt.Println(names)

	//map声明2 直接make
	var names2 = make(map[int]string)
	names2[0] = "Go"
	names2[1] = "C"
	names2[2] = "Python"
	names2[3] = "Java"
	fmt.Println(names2)

	//map声明3 直接赋值
	var names3 map[int]string = map[int]string{
		0: "Go",
		1: "C",
		2: "Python",
		3: "Java"}

	fmt.Println(names3)

	delete(names3, 0)
	fmt.Println(names3)
	//- 1.删除所有的key，没有专门的方法（类似map.clear()），可以遍历一下key，逐个删除
	//- 2.或者map=make(...) ,make一个新的 让原来的成为垃圾被gc回收
	names3 = make(map[int]string)
	fmt.Println(names3)
	s, isFind := names2[12]
	if isFind {
		fmt.Println(s)
	} else {
		fmt.Println("没找到")
	}
	//map遍历
	for i, value := range names {
		fmt.Printf("key: %d value：%s \n", i, value)
	}

	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "illusory"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "重庆"

	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "Azz"
	studentMap["stu02"]["sex"] = "男"
	studentMap["stu02"]["address"] = "成都"

	for i, value := range studentMap {
		fmt.Printf("key: %s value：%s \n", i, value)
		for j, value2 := range value {
			fmt.Printf("key: %s value：%s \n", j, value2)
		}
	}
	//map切片
	var userMap []map[string]string
	//切片需要make
	userMap = make([]map[string]string, 2)
	//map也需要make
	userMap[0] = make(map[string]string, 2)
	userMap[0]["name"] = "illusory"
	userMap[0]["age"] = "22"
	userMap[1] = make(map[string]string, 2)
	userMap[1]["name"] = "Azz"
	userMap[1]["age"] = "22"
	//越界了 切片动态扩容 使用append
	//userMap[2] = make(map[string]string, 2)
	//userMap[2]["name"] = "webpack"
	//userMap[2]["age"] = "22"

	newUser := map[string]string{
		"name": "newUser",
		"age":  "30"}

	userMap = append(userMap, newUser)
	fmt.Println(userMap)

	sortMap := make(map[int]int, 10)
	sortMap[10] = 11
	sortMap[20] = 39
	sortMap[8] = 30
	sortMap[1] = 12
	fmt.Println(sortMap)

	//按照map的key的顺序排序输出
	//1.先将map的key放入切片
	//2.对切片排序
	//3.遍历切片 按照key输出map的值
	var keys []int
	for i, _ := range sortMap {
		keys = append(keys, i)
	}
	//排序sort.Ints()
	sort.Ints(keys)
	fmt.Println(keys)
	//遍历切片 输出value
	for _, key := range keys {
		fmt.Printf("key %d value %d \t", key, sortMap[key])
	}
}
