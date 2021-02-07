package lcof

import "strings"

// https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
func replaceSpace(s string) string {
	var sb strings.Builder
	for _, v := range s {
		if v == ' ' {
			sb.WriteString("%20")
		} else {
			sb.WriteRune(v)
		}
	}
	return sb.String()
}

func replaceSpaceStand(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

func replaceSpace2(s string) string {
	str := make([]rune, 0, len(s)*3)
	for _, v := range s {
		if v == ' ' {
			str = append(str, '%', '2', '0')
		} else {
			str = append(str, v)
		}
	}
	return string(str)
}

func replaceSpace3(s string) string {
	result := make([]rune, len(s)*3)

	i := 0
	for _, v := range s {
		if v != ' ' {
			result[i] = v
			i++
		} else {
			result[i] = '%'
			result[i+1] = '2'
			result[i+2] = '0'
			i += 3
		}
	}

	return string(result)[:i]
}
