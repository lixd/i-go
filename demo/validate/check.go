package validate

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode"
)

// 电话校验
func Phone(phone string) bool {
	matched, err := regexp.MatchString("^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\\d{8}$", phone)
	if err != nil {
		fmt.Println(err)
	}
	return matched
}

// IsNumbers 字符串是否全为数字
func IsNumbers(keyWords string) bool {
	str2 := []rune(keyWords)
	for _, v := range str2 {
		isNumber := unicode.IsNumber(v)
		if !isNumber {
			return false
		}
	}
	return true
}

// ParseItemId 从 URL 中解析出itemId
func ParseItemId(URL string) (string, error) {
	var itemId string
	parse, err := url.Parse(URL)
	if err != nil {
		return "", err
	}
	values, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return "", err
	}
	if len(values) != 0 && len(values["id"]) != 0 {
		itemId = values["id"][0]
	}
	return itemId, nil
}

// IsURL 是否为 URL
func IsURL(URL string) bool {
	// 简单判断一下 包含 http 就算是 URL
	isContains := strings.Contains(URL, "http")
	return isContains
}
