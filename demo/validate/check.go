package validate

import (
	"fmt"
	"regexp"
	"unicode"
)

// Phone 电话校验
func Phone(phone string) bool {
	matched, err := regexp.MatchString("^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\\d{8}$", phone)
	if err != nil {
		fmt.Println(err)
	}
	return matched
}

// IsNumbers 字符串是否全为数字
func IsNumbers(keyWords string) bool {
	for _, v := range keyWords {
		isNumber := unicode.IsNumber(v)
		if !isNumber {
			return false
		}
	}
	return true
}
