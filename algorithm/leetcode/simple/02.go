package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
面试题 01.02. 判定是否互为字符重排
给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

示例 1：

输入: s1 = "abc", s2 = "bca"
输出: true
示例 2：

输入: s1 = "abc", s2 = "bad"
输出: false
说明：

0 <= len(s1) <= 100
0 <= len(s2) <= 100
通过次数12,212提交次数18,258

*/
func main() {
	s1 := "krqdpwdvgfuogtobtyylexrebrwzynzlpkotoqmokfpqeibbhzdjcwpgprzpqersmmdxdmwssfbfwmmvrxkjyjteirloxpbiopso"
	s2 := "pyymgxtdqzqxxkmirptmbewjobpslwkbmmzfbwzmltowevsofkydrejdpcoripjlewoqzgusvypotrdkepbqspxdmoyrfnyrbrof"
	fmt.Println(CheckPermutation(s1, s2))
}

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := make(map[int32]int)
	for _, v := range s1 {
		m[v] += 1
	}
	for _, v := range s2 {
		if _, ok := m[v]; ok {
			m[v]--
		} else {
			return false
		}
	}
	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}

/*
首先将s1,s2利用strings.Split()函数拆分成单个字母的字符串切片，然后再利用sort.Strings()函数将字符串切片进行排序，
最后利用strings.Join()函数将排序后的字符串切片组合成字符串，最后返回最终的两个字符串是否相等就能判断另一个字符串是否是重排列得到的
*/
func CheckPermutation2(s1 string, s2 string) bool {
	tmp1 := strings.Split(s1, "")
	tmp2 := strings.Split(s2, "")
	sort.Strings(tmp1)
	res1 := strings.Join(tmp1, "")
	sort.Strings(tmp2)
	res2 := strings.Join(tmp2, "")
	return res1 == res2
}
