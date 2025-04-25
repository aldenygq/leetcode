//题目要求
//输入一个字符串，检测字符串中是否存在连续重复字符，如存在，则计算重复字符数量，替代重复字符。
//案例："AAbcCCDdefGsffffrvyyuhkkkkdf"输出"A2bcC2DdefGsf4rvy2uhk4df"


package main

import (
	"fmt"
)

func main() {
	str := "AAbcCCDdefGsffffrvyyuhkkkkdf"
	result := compressString(str)
	fmt.Println(result)
}

func compressString(s string) string {
	if len(s) == 0 {
		return ""
	}

	var result []rune
	currentChar := rune(s[0])
	count := 1

	for i := 1; i < len(s); i++ {
		if rune(s[i]) == currentChar {
			count++
		} else {
			result = append(result, currentChar)
			if count > 1 {
				result = append(result, []rune(fmt.Sprint(count))...)
			}
			currentChar = rune(s[i])
			count = 1
		}
	}

	// 添加最后一个字符组
	result = append(result, currentChar)
	if count > 1 {
		result = append(result, []rune(fmt.Sprint(count))...)
	}

	return string(result)
}
